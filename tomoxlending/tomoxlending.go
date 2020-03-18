package tomoxlending

import (
	"errors"
	"fmt"
	"github.com/tomochain/tomochain/consensus"
	"github.com/tomochain/tomochain/core/types"
	"github.com/tomochain/tomochain/p2p"
	"github.com/tomochain/tomochain/tomox"
	"github.com/tomochain/tomochain/tomox/tradingstate"
	"github.com/tomochain/tomochain/tomoxDAO"
	"github.com/tomochain/tomochain/tomoxlending/lendingstate"
	"gopkg.in/karalabe/cookiejar.v2/collections/prque"
	"math/big"
	"strconv"
	"time"

	lru "github.com/hashicorp/golang-lru"
	"github.com/tomochain/tomochain/common"
	"github.com/tomochain/tomochain/core/state"
	"github.com/tomochain/tomochain/log"
	"github.com/tomochain/tomochain/rpc"
)

const (
	ProtocolName       = "tomoxlending"
	ProtocolVersion    = uint64(1)
	ProtocolVersionStr = "1.0"
	defaultCacheLimit  = 1024
)

var (
	ErrNonceTooHigh = errors.New("nonce too high")
	ErrNonceTooLow  = errors.New("nonce too low")
)

type Lending struct {
	Triegc     *prque.Prque          // Priority queue mapping block numbers to tries to gc
	StateCache lendingstate.Database // State database to reuse between imports (contains state cache)    *lendingstate.TradingStateDB

	orderNonce map[common.Address]*big.Int

	tomox               *tomox.TomoX
	lendingItemHistory  *lru.Cache
	lendingTradeHistory *lru.Cache
}

func (l *Lending) Protocols() []p2p.Protocol {
	return []p2p.Protocol{}
}

func (l *Lending) Start(server *p2p.Server) error {
	return nil
}

func (l *Lending) Stop() error {
	return nil
}

func New(tomox *tomox.TomoX) *Lending {
	itemCache, _ := lru.New(defaultCacheLimit)
	lendingTradeCache, _ := lru.New(defaultCacheLimit)
	lending := &Lending{
		orderNonce:          make(map[common.Address]*big.Int),
		Triegc:              prque.New(),
		lendingItemHistory:  itemCache,
		lendingTradeHistory: lendingTradeCache,
	}
	lending.StateCache = lendingstate.NewDatabase(tomox.GetLevelDB())
	lending.tomox = tomox
	return lending
}

func (l *Lending) GetLevelDB() tomoxDAO.TomoXDAO {
	return l.tomox.GetLevelDB()
}

func (l *Lending) GetMongoDB() tomoxDAO.TomoXDAO {
	return l.tomox.GetMongoDB()
}

// APIs returns the RPC descriptors the Lending implementation offers
func (l *Lending) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: ProtocolName,
			Version:   ProtocolVersionStr,
			Service:   NewPublicTomoXLendingAPI(l),
			Public:    true,
		},
	}
}

// Version returns the Lending sub-protocols version number.
func (l *Lending) Version() uint64 {
	return ProtocolVersion
}

func (l *Lending) ProcessOrderPending(createdBlockTime uint64, coinbase common.Address, chain consensus.ChainContext, pending map[common.Address]types.LendingTransactions, statedb *state.StateDB, lendingStatedb *lendingstate.LendingStateDB, tradingStateDb *tradingstate.TradingStateDB) ([]*lendingstate.LendingItem, map[common.Hash]lendingstate.MatchingResult) {
	lendingItems := []*lendingstate.LendingItem{}
	matchingResults := map[common.Hash]lendingstate.MatchingResult{}

	txs := types.NewLendingTransactionByNonce(types.LendingTxSigner{}, pending)
	for {
		tx := txs.Peek()
		if tx == nil {
			break
		}
		log.Debug("ProcessOrderPending start", "len", len(pending))
		log.Debug("Get pending orders to process", "address", tx.UserAddress(), "nonce", tx.Nonce())
		V, R, S := tx.Signature()

		bigstr := V.String()
		n, e := strconv.ParseInt(bigstr, 10, 8)
		if e != nil {
			continue
		}

		order := &lendingstate.LendingItem{
			Nonce:           big.NewInt(int64(tx.Nonce())),
			Quantity:        tx.Quantity(),
			Interest:        new(big.Int).SetUint64(tx.Interest()),
			Relayer:         tx.RelayerAddress(),
			Term:            tx.Term(),
			UserAddress:     tx.UserAddress(),
			LendingToken:    tx.LendingToken(),
			CollateralToken: tx.CollateralToken(),
			AutoTopUp:       tx.AutoTopUp(),
			Status:          tx.Status(),
			Side:            tx.Side(),
			Type:            tx.Type(),
			Hash:            tx.LendingHash(),
			LendingId:       tx.LendingId(),
			LendingTradeId:  tx.LendingTradeId(),
			ExtraData:       tx.ExtraData(),
			Signature: &lendingstate.Signature{
				V: byte(n),
				R: common.BigToHash(R),
				S: common.BigToHash(S),
			},
		}
		cancel := false
		if order.Status == lendingstate.LendingStatusCancelled {
			cancel = true
		}

		log.Info("Process order pending", "orderPending", order, "LendingToken", order.LendingToken.Hex(), "CollateralToken", order.CollateralToken)
		originalOrder := &lendingstate.LendingItem{}
		*originalOrder = *order
		originalOrder.Quantity = lendingstate.CloneBigInt(order.Quantity)

		if cancel {
			order.Status = lendingstate.LendingStatusCancelled
		}

		newTrades, newRejectedOrders, err := l.CommitOrder(createdBlockTime, coinbase, chain, statedb, lendingStatedb, tradingStateDb, lendingstate.GetLendingOrderBookHash(order.LendingToken, order.Term), order)
		for _, reject := range newRejectedOrders {
			log.Debug("Reject order", "reject", *reject)
		}

		switch err {
		case ErrNonceTooLow:
			// New head notification data race between the transaction pool and miner, shift
			log.Debug("Skipping order with low nonce", "sender", tx.UserAddress(), "nonce", tx.Nonce())
			txs.Shift()
			continue

		case ErrNonceTooHigh:
			// Reorg notification data race between the transaction pool and miner, skip account =
			log.Debug("Skipping order account with high nonce", "sender", tx.UserAddress(), "nonce", tx.Nonce())
			txs.Pop()
			continue

		case nil:
			// everything ok
			txs.Shift()

		default:
			// Strange error, discard the transaction and get the next in line (note, the
			// nonce-too-high clause will prevent us from executing in vain).
			log.Debug("Transaction failed, account skipped", "hash", tx.Hash(), "err", err)
			txs.Shift()
			continue
		}

		// orderID has been updated
		originalOrder.LendingId = order.LendingId
		lendingItems = append(lendingItems, originalOrder)
		matchingResults[order.Hash] = lendingstate.MatchingResult{
			Trades:  newTrades,
			Rejects: newRejectedOrders,
		}
	}
	return lendingItems, matchingResults
}

// there are 3 tasks need to complete (for SDK nodes) after matching
// 1. Put takerLendingItem to database
// 2.a Update status, filledAmount of makerLendingItem
// 2.b. Put lendingTrade to database
// 3. Update status of rejected items
func (l *Lending) SyncDataToSDKNode(takerLendingItem *lendingstate.LendingItem, txHash common.Hash, txMatchTime time.Time, trades []*lendingstate.LendingTrade, rejectedItems []*lendingstate.LendingItem, dirtyOrderCount *uint64) error {
	var (
		// originTakerLendingItem: item getting from database
		originTakerLendingItem, updatedTakerLendingItem *lendingstate.LendingItem
		makerDirtyHashes                                []string
		makerDirtyFilledAmount                          map[string]*big.Int
		err                                             error
	)
	db := l.GetMongoDB()
	db.InitLendingBulk()
	// 1. put processed takerLendingItem to database
	lastState := lendingstate.LendingItemHistoryItem{}
	// Typically, takerItem has never existed in database
	// except cancel case: in this case, item existed in database with status = OPEN, then use send another lendingItem to cancel it
	val, err := db.GetObject(takerLendingItem.Hash, &lendingstate.LendingItem{})
	if err == nil && val != nil {
		originTakerLendingItem = val.(*lendingstate.LendingItem)
		lastState = lendingstate.LendingItemHistoryItem{
			TxHash:       originTakerLendingItem.TxHash,
			FilledAmount: lendingstate.CloneBigInt(originTakerLendingItem.FilledAmount),
			Status:       originTakerLendingItem.Status,
			UpdatedAt:    originTakerLendingItem.UpdatedAt,
		}
	}
	if originTakerLendingItem != nil {
		updatedTakerLendingItem = originTakerLendingItem
	} else {
		updatedTakerLendingItem = takerLendingItem
		updatedTakerLendingItem.FilledAmount = new(big.Int)
	}

	if takerLendingItem.Status == lendingstate.LendingStatusNew {
		updatedTakerLendingItem.Status = lendingstate.LendingStatusOpen
	} else if takerLendingItem.Status == lendingstate.LendingStatusCancelled {
		updatedTakerLendingItem.Status = lendingstate.LendingStatusCancelled
	}

	updatedTakerLendingItem.TxHash = txHash
	if updatedTakerLendingItem.CreatedAt.IsZero() {
		updatedTakerLendingItem.CreatedAt = txMatchTime
	}
	if txMatchTime.Before(updatedTakerLendingItem.UpdatedAt) || (txMatchTime.Equal(updatedTakerLendingItem.UpdatedAt) && *dirtyOrderCount == 0) {
		log.Debug("Ignore old lendingItem/lendingTrades taker", "txHash", txHash.Hex(), "txTime", txMatchTime.UnixNano(), "updatedAt", updatedTakerLendingItem.UpdatedAt.UnixNano())
		return nil
	}
	*dirtyOrderCount++

	l.UpdateLendingItemCache(updatedTakerLendingItem.LendingToken, updatedTakerLendingItem.CollateralToken, updatedTakerLendingItem.Hash, txHash, lastState)
	updatedTakerLendingItem.UpdatedAt = txMatchTime

	// 2. put trades to database and update status
	log.Debug("Got lendingTrades", "number", len(trades), "txhash", txHash.Hex())
	makerDirtyFilledAmount = make(map[string]*big.Int)

	tradeList := map[common.Hash]*lendingstate.LendingTrade{}
	for _, tradeRecord := range trades {
		// 2.a. put to trades
		if tradeRecord.Status != lendingstate.TradeStatusOpen {
			log.Debug("UpdateLendingTrade:", "hash", tradeRecord.Hash.Hex(), "status", tradeRecord.Status, "tradeId", tradeRecord.TradeId)
			tradeList[tradeRecord.Hash] = tradeRecord
			//if err := l.UpdateLendingTrade([]common.Hash{tradeRecord.Hash}, tradeRecord.Status, txHash, txMatchTime); err != nil {
			//	return err
			//}
			continue
		}
		if tradeRecord.CreatedAt.IsZero() {
			tradeRecord.CreatedAt = txMatchTime
		}
		tradeRecord.UpdatedAt = txMatchTime
		tradeRecord.TxHash = txHash
		tradeRecord.Hash = tradeRecord.ComputeHash()
		tradeList[tradeRecord.Hash] = tradeRecord

		// 2.b. update status and filledAmount
		filledAmount := new(big.Int)
		if tradeRecord.Amount != nil {
			filledAmount = lendingstate.CloneBigInt(tradeRecord.Amount)
		}
		// maker dirty order
		makerFilledAmount := big.NewInt(0)
		makerOrderHash := common.Hash{}
		if updatedTakerLendingItem.Side == lendingstate.Borrowing {
			makerOrderHash = tradeRecord.InvestingOrderHash
		} else {
			makerOrderHash = tradeRecord.BorrowingOrderHash
		}
		if amount, ok := makerDirtyFilledAmount[makerOrderHash.Hex()]; ok {
			makerFilledAmount = lendingstate.CloneBigInt(amount)
		}
		makerFilledAmount = new(big.Int).Add(makerFilledAmount, filledAmount)
		makerDirtyFilledAmount[makerOrderHash.Hex()] = makerFilledAmount
		makerDirtyHashes = append(makerDirtyHashes, makerOrderHash.Hex())

		if updatedTakerLendingItem.Status != lendingstate.Repay && updatedTakerLendingItem.Status != lendingstate.TopUp {
			//updatedTakerOrder = l.updateMatchedOrder(updatedTakerOrder, filledAmount, txMatchTime, txHash)
			//  update filledAmount, status of takerOrder
			updatedTakerLendingItem.FilledAmount = new(big.Int).Add(updatedTakerLendingItem.FilledAmount, filledAmount)
			if updatedTakerLendingItem.FilledAmount.Cmp(updatedTakerLendingItem.Quantity) < 0 && updatedTakerLendingItem.Type == lendingstate.Limit {
				updatedTakerLendingItem.Status = lendingstate.LendingStatusPartialFilled
			} else {
				updatedTakerLendingItem.Status = lendingstate.LendingStatusFilled
			}
		}
	}
	if err := l.UpdateLendingTrade(tradeList, txHash, txMatchTime); err != nil {
		return err
	}

	// update status for Market orders
	if updatedTakerLendingItem.Type == lendingstate.Market && updatedTakerLendingItem.Status != lendingstate.Repay && updatedTakerLendingItem.Status != lendingstate.TopUp {
		if updatedTakerLendingItem.FilledAmount.Cmp(big.NewInt(0)) > 0 {
			updatedTakerLendingItem.Status = lendingstate.LendingStatusFilled
		} else {
			updatedTakerLendingItem.Status = lendingstate.LendingStatusReject
		}
	}

	log.Debug("PutObject processed takerLendingItem",
		"term", updatedTakerLendingItem.Term, "userAddr", updatedTakerLendingItem.UserAddress.Hex(), "side", updatedTakerLendingItem.Side,
		"Interest", updatedTakerLendingItem.Interest, "quantity", updatedTakerLendingItem.Quantity, "filledAmount", updatedTakerLendingItem.FilledAmount, "status", updatedTakerLendingItem.Status,
		"hash", updatedTakerLendingItem.Hash.Hex(), "txHash", updatedTakerLendingItem.TxHash.Hex())

	if err := db.PutObject(updatedTakerLendingItem.Hash, updatedTakerLendingItem); err != nil {
		return fmt.Errorf("SDKNode: failed to put processed takerOrder. Hash: %s Error: %s", updatedTakerLendingItem.Hash.Hex(), err.Error())
	}

	items := db.GetListItemByHashes(makerDirtyHashes, &lendingstate.LendingItem{})
	if items != nil {
		makerItems := items.([]*lendingstate.LendingItem)
		log.Debug("Maker dirty lendingItem", "len", len(makerItems), "txhash", txHash.Hex())
		for _, m := range makerItems {
			if txMatchTime.Before(m.UpdatedAt) {
				log.Debug("Ignore old lendingItem/lendingTrades maker", "txHash", txHash.Hex(), "txTime", txMatchTime.UnixNano(), "updatedAt", m.UpdatedAt.UnixNano())
				continue
			}
			lastState = lendingstate.LendingItemHistoryItem{
				TxHash:       m.TxHash,
				FilledAmount: lendingstate.CloneBigInt(m.FilledAmount),
				Status:       m.Status,
				UpdatedAt:    m.UpdatedAt,
			}
			l.UpdateLendingItemCache(m.LendingToken, m.CollateralToken, m.Hash, txHash, lastState)
			m.TxHash = txHash
			m.UpdatedAt = txMatchTime
			m.FilledAmount = new(big.Int).Add(m.FilledAmount, makerDirtyFilledAmount[m.Hash.Hex()])
			if m.FilledAmount.Cmp(m.Quantity) < 0 {
				m.Status = lendingstate.LendingStatusPartialFilled
			} else {
				m.Status = lendingstate.LendingStatusFilled
			}
			log.Debug("PutObject processed makerLendingItem",
				"term", m.Term, "userAddr", m.UserAddress.Hex(), "side", m.Side,
				"Interest", m.Interest, "quantity", m.Quantity, "filledAmount", m.FilledAmount, "status", m.Status,
				"hash", m.Hash.Hex(), "txHash", m.TxHash.Hex())
			if err := db.PutObject(m.Hash, m); err != nil {
				return fmt.Errorf("SDKNode: failed to put processed makerOrder. Hash: %s Error: %s", m.Hash.Hex(), err.Error())
			}
		}
	}

	// 3. put rejected orders to leveldb and update status REJECTED
	log.Debug("Got rejected lendingItems", "number", len(rejectedItems), "rejectedLendingItems", rejectedItems)

	if len(rejectedItems) > 0 {
		var rejectedHashes []string
		// updateRejectedOrders
		for _, r := range rejectedItems {
			rejectedHashes = append(rejectedHashes, r.Hash.Hex())
			if updatedTakerLendingItem.Hash == r.Hash && !txMatchTime.Before(r.UpdatedAt) {
				// cache r history for handling reorg
				historyRecord := lendingstate.LendingItemHistoryItem{
					TxHash:       updatedTakerLendingItem.TxHash,
					FilledAmount: lendingstate.CloneBigInt(updatedTakerLendingItem.FilledAmount),
					Status:       updatedTakerLendingItem.Status,
					UpdatedAt:    updatedTakerLendingItem.UpdatedAt,
				}
				l.UpdateLendingItemCache(updatedTakerLendingItem.LendingToken, updatedTakerLendingItem.CollateralToken, updatedTakerLendingItem.Hash, txHash, historyRecord)

				updatedTakerLendingItem.Status = lendingstate.LendingStatusReject
				updatedTakerLendingItem.TxHash = txHash
				updatedTakerLendingItem.UpdatedAt = txMatchTime
				if err := db.PutObject(updatedTakerLendingItem.Hash, updatedTakerLendingItem); err != nil {
					return fmt.Errorf("SDKNode: failed to reject takerOrder. Hash: %s Error: %s", updatedTakerLendingItem.Hash.Hex(), err.Error())
				}
			}
		}
		items := db.GetListItemByHashes(rejectedHashes, &lendingstate.LendingItem{})
		if items != nil {
			dirtyRejectedItems := items.([]*lendingstate.LendingItem)
			for _, r := range dirtyRejectedItems {
				if txMatchTime.Before(r.UpdatedAt) {
					log.Debug("Ignore old orders/trades reject", "txHash", txHash.Hex(), "txTime", txMatchTime.UnixNano(), "updatedAt", updatedTakerLendingItem.UpdatedAt.UnixNano())
					continue
				}
				// cache lendingItem for handling reorg
				historyRecord := lendingstate.LendingItemHistoryItem{
					TxHash:       r.TxHash,
					FilledAmount: lendingstate.CloneBigInt(r.FilledAmount),
					Status:       r.Status,
					UpdatedAt:    r.UpdatedAt,
				}
				l.UpdateLendingItemCache(r.LendingToken, r.CollateralToken, r.Hash, txHash, historyRecord)
				dirtyFilledAmount, ok := makerDirtyFilledAmount[r.Hash.Hex()]
				if ok && dirtyFilledAmount != nil {
					r.FilledAmount = new(big.Int).Add(r.FilledAmount, dirtyFilledAmount)
				}
				r.Status = lendingstate.LendingStatusReject
				r.TxHash = txHash
				r.UpdatedAt = txMatchTime
				if err = db.PutObject(r.Hash, r); err != nil {
					return fmt.Errorf("SDKNode: failed to update rejectedOder to sdkNode %s", err.Error())
				}
			}
		}
	}

	if err := db.CommitLendingBulk(); err != nil {
		return fmt.Errorf("SDKNode fail to commit bulk update lendingItem/lendingTrades at txhash %s . Error: %s", txHash.Hex(), err.Error())
	}
	return nil
}

func (l *Lending) UpdateLiquidatedTrade(result lendingstate.FinalizedResult, trades map[common.Hash]*lendingstate.LendingTrade) error {
	db := l.GetMongoDB()
	db.InitLendingBulk()


	txhash := result.TxHash
	txTime := time.Unix(0, (result.Timestamp/1e6)*1e6).UTC() // round to milliseconds
	if err := l.UpdateLendingTrade(trades, txhash, txTime); err != nil {
		return err
	}

	// adding auto repay transaction
	if len(result.AutoRepay) > 0 {
		for _, hash := range result.AutoRepay {
			trade := trades[hash]
			if trade == nil {
				continue
			}
			repayItem := &lendingstate.LendingItem{
				Quantity:        trade.Amount,
				Interest:        big.NewInt(int64(trade.Interest)),
				Side:            "",
				Type:            "",
				LendingToken:    trade.LendingToken,
				CollateralToken: trade.CollateralToken,
				FilledAmount:    nil,
				Status:          lendingstate.Repay,
				Relayer:         trade.BorrowingRelayer,
				Term:            trade.Term,
				UserAddress:     trade.Borrower,
				Signature:       nil,
				Hash:            trade.BorrowingOrderHash,
				TxHash:          txhash,
				Nonce:           nil,
				CreatedAt:       txTime,
				UpdatedAt:       txTime,
				LendingId:       0,
				LendingTradeId:  trade.TradeId,
				ExtraData:       "auto",
			}
			if err := db.PutObject(repayItem.Hash, repayItem); err != nil {
				return err
			}
		}
	}

	// adding auto topup transaction
	if len(result.AutoTopUp) > 0 {
		oldTradeHashes := []string{}
		for _, hash := range result.AutoTopUp {
			oldTradeHashes = append(oldTradeHashes, hash.Hex())
		}
		items := db.GetListItemByHashes(oldTradeHashes, &lendingstate.LendingTrade{})
		if items != nil && len(items.([]*lendingstate.LendingTrade)) > 0 {
			for _, oldTrade := range items.([]*lendingstate.LendingTrade) {
				newTrade := trades[oldTrade.Hash]
				topUpItem := &lendingstate.LendingItem{
					Quantity:        new(big.Int).Sub(newTrade.CollateralLockedAmount, oldTrade.CollateralLockedAmount),
					Interest:        big.NewInt(int64(oldTrade.Interest)),
					Side:            "",
					Type:            "",
					LendingToken:    oldTrade.LendingToken,
					CollateralToken: oldTrade.CollateralToken,
					FilledAmount:    nil,
					Status:          lendingstate.TopUp,
					AutoTopUp:       true,
					Relayer:         oldTrade.BorrowingRelayer,
					Term:            oldTrade.Term,
					UserAddress:     oldTrade.Borrower,
					Signature:       nil,
					Hash:            oldTrade.BorrowingOrderHash,
					TxHash:          txhash,
					Nonce:           nil,
					CreatedAt:       txTime,
					UpdatedAt:       txTime,
					LendingId:       0,
					LendingTradeId:  oldTrade.TradeId,
					ExtraData:       "auto",
				}
				if err := db.PutObject(topUpItem.Hash, topUpItem); err != nil {
					return err
				}
			}
		}
	}

	if err := db.CommitLendingBulk(); err != nil {
		return fmt.Errorf("failed to updateLendingTrade . Err: %v", err)
	}

	return nil
}

func (l *Lending) UpdateLendingTrade(trades map[common.Hash]*lendingstate.LendingTrade, txhash common.Hash, txTime time.Time) error {
	db := l.GetMongoDB()
	hashQuery := []string{}
	if len(trades) == 0 {
		return nil
	}
	for _, trade := range trades {
		hashQuery = append(hashQuery, trade.Hash.Hex())
	}
	items := db.GetListItemByHashes(hashQuery, &lendingstate.LendingTrade{})
	if items != nil && len(items.([]*lendingstate.LendingTrade)) > 0 {
		for _, trade := range items.([]*lendingstate.LendingTrade) {
			history := lendingstate.LendingTradeHistoryItem{
				TxHash:                 trade.TxHash,
				CollateralLockedAmount: trade.CollateralLockedAmount,
				LiquidationPrice:       trade.LiquidationPrice,
				Status:                 trade.Status,
				UpdatedAt:              trade.UpdatedAt,
			}
			l.UpdateLendingTradeCache(trade.Hash, txhash, history)
			trade.TxHash = txhash
			trade.UpdatedAt = txTime

			newAutoTopUpTrade := trades[trade.Hash]
			trade.CollateralLockedAmount = newAutoTopUpTrade.CollateralLockedAmount
			trade.Status = newAutoTopUpTrade.Status
			trade.LiquidationPrice = newAutoTopUpTrade.LiquidationPrice

			if err := db.PutObject(trade.Hash, trade); err != nil {
				return err
			}
		}
		log.Debug("UpdateLendingTrade successfully", "txhash", txhash, "hash", hashQuery)
	} else {
		// not update, just upsert
		for _, trade := range trades {
			if err := db.PutObject(trade.Hash, trade); err != nil {
				return err
			}
		}
	}
	return nil
}

func (l *Lending) GetLendingState(block *types.Block) (*lendingstate.LendingStateDB, error) {
	root, err := l.GetLendingStateRoot(block)
	if err != nil {
		return nil, err
	}
	if l.StateCache == nil {
		return nil, errors.New("Not initialized tomox")
	}
	return lendingstate.New(root, l.StateCache)
}

func (l *Lending) GetStateCache() lendingstate.Database {
	return l.StateCache
}

func (l *Lending) GetTriegc() *prque.Prque {
	return l.Triegc
}

func (l *Lending) GetLendingStateRoot(block *types.Block) (common.Hash, error) {
	for _, tx := range block.Transactions() {
		if tx.To() != nil && tx.To().Hex() == common.TradingStateAddr {
			if len(tx.Data()) >= 64 {
				return common.BytesToHash(tx.Data()[32:]), nil
			}
		}
	}
	return lendingstate.EmptyRoot, nil
}

func (l *Lending) UpdateLendingItemCache(LendingToken, CollateralToken common.Address, hash common.Hash, txhash common.Hash, lastState lendingstate.LendingItemHistoryItem) {
	var lendingCacheAtTxHash map[common.Hash]lendingstate.LendingItemHistoryItem
	c, ok := l.lendingItemHistory.Get(txhash)
	if !ok || c == nil {
		lendingCacheAtTxHash = make(map[common.Hash]lendingstate.LendingItemHistoryItem)
	} else {
		lendingCacheAtTxHash = c.(map[common.Hash]lendingstate.LendingItemHistoryItem)
	}
	orderKey := lendingstate.GetLendingItemHistoryKey(LendingToken, CollateralToken, hash)
	_, ok = lendingCacheAtTxHash[orderKey]
	if !ok {
		lendingCacheAtTxHash[orderKey] = lastState
	}
	l.lendingItemHistory.Add(txhash, lendingCacheAtTxHash)
}

func (l *Lending) UpdateLendingTradeCache(hash common.Hash, txhash common.Hash, lastState lendingstate.LendingTradeHistoryItem) {
	var lendingCacheAtTxHash map[common.Hash]lendingstate.LendingTradeHistoryItem
	c, ok := l.lendingTradeHistory.Get(txhash)
	if !ok || c == nil {
		lendingCacheAtTxHash = make(map[common.Hash]lendingstate.LendingTradeHistoryItem)
	} else {
		lendingCacheAtTxHash = c.(map[common.Hash]lendingstate.LendingTradeHistoryItem)
	}
	_, ok = lendingCacheAtTxHash[hash]
	if !ok {
		lendingCacheAtTxHash[hash] = lastState
	}
	l.lendingTradeHistory.Add(txhash, lendingCacheAtTxHash)
}

func (l *Lending) RollbackLendingData(txhash common.Hash) error {
	db := l.GetMongoDB()
	db.InitLendingBulk()

	// rollback lendingItem
	items := db.GetListItemByTxHash(txhash, &lendingstate.LendingItem{})
	if items != nil {
		for _, item := range items.([]*lendingstate.LendingItem) {
			c, ok := l.lendingItemHistory.Get(txhash)
			log.Debug("tomoxlending reorg: rollback lendingItem", "txhash", txhash.Hex(), "item", lendingstate.ToJSON(item), "lendingItemHistory", c)
			if !ok {
				log.Debug("tomoxlending reorg: remove item due to no lendingItemHistory", "item", lendingstate.ToJSON(item))
				if err := db.DeleteObject(item.Hash, &lendingstate.LendingItem{}); err != nil {
					return fmt.Errorf("failed to remove reorg LendingItem. Err: %v . Item: %s", err.Error(), lendingstate.ToJSON(item))
				}
				continue
			}
			cacheAtTxHash := c.(map[common.Hash]lendingstate.LendingItemHistoryItem)
			lendingItemHistory, _ := cacheAtTxHash[lendingstate.GetLendingItemHistoryKey(item.LendingToken, item.CollateralToken, item.Hash)]
			if (lendingItemHistory == lendingstate.LendingItemHistoryItem{}) {
				log.Debug("tomoxlending reorg: remove item due to empty lendingItemHistory", "item", lendingstate.ToJSON(item))
				if err := db.DeleteObject(item.Hash, &lendingstate.LendingItem{}); err != nil {
					return fmt.Errorf("failed to remove reorg LendingItem. Err: %v . Item: %s", err.Error(), lendingstate.ToJSON(item))
				}
				continue
			}
			item.TxHash = lendingItemHistory.TxHash
			item.Status = lendingItemHistory.Status
			item.FilledAmount = lendingstate.CloneBigInt(lendingItemHistory.FilledAmount)
			item.UpdatedAt = lendingItemHistory.UpdatedAt
			log.Debug("tomoxlending reorg: update item to the last lendingItemHistory", "item", lendingstate.ToJSON(item), "lendingItemHistory", lendingItemHistory)
			if err := db.PutObject(item.Hash, item); err != nil {
				return fmt.Errorf("failed to update reorg LendingItem. Err: %v . Item: %s", err.Error(), lendingstate.ToJSON(item))
			}
		}
	}

	// rollback lendingTrade
	items = db.GetListItemByTxHash(txhash, &lendingstate.LendingTrade{})
	if items != nil {
		for _, trade := range items.([]*lendingstate.LendingTrade) {
			c, ok := l.lendingTradeHistory.Get(txhash)
			log.Debug("tomoxlending reorg: rollback LendingTrade", "txhash", txhash.Hex(), "trade", lendingstate.ToJSON(trade), "LendingTradeHistory", c)
			if !ok {
				log.Debug("tomoxlending reorg: remove trade due to no LendingTradeHistory", "trade", lendingstate.ToJSON(trade))
				if err := db.DeleteObject(trade.Hash, &lendingstate.LendingTrade{}); err != nil {
					return fmt.Errorf("failed to remove reorg LendingTrade. Err: %v . Trade: %s", err.Error(), lendingstate.ToJSON(trade))
				}
				continue
			}
			cacheAtTxHash := c.(map[common.Hash]lendingstate.LendingTradeHistoryItem)
			lendingTradeHistoryItem, _ := cacheAtTxHash[trade.Hash]
			if (lendingTradeHistoryItem == lendingstate.LendingTradeHistoryItem{}) {
				log.Debug("tomoxlending reorg: remove trade due to empty LendingTradeHistory", "trade", lendingstate.ToJSON(trade))
				if err := db.DeleteObject(trade.Hash, &lendingstate.LendingTrade{}); err != nil {
					return fmt.Errorf("failed to remove reorg LendingTrade. Err: %v . Trade: %s", err.Error(), lendingstate.ToJSON(trade))
				}
				continue
			}
			trade.TxHash = lendingTradeHistoryItem.TxHash
			trade.Status = lendingTradeHistoryItem.Status
			trade.CollateralLockedAmount = lendingstate.CloneBigInt(lendingTradeHistoryItem.CollateralLockedAmount)
			trade.LiquidationPrice = lendingstate.CloneBigInt(lendingTradeHistoryItem.LiquidationPrice)
			trade.UpdatedAt = lendingTradeHistoryItem.UpdatedAt
			log.Debug("tomoxlending reorg: update trade to the last lendingTradeHistoryItem", "trade", lendingstate.ToJSON(trade), "lendingTradeHistoryItem", lendingTradeHistoryItem)
			if err := db.PutObject(trade.Hash, trade); err != nil {
				return fmt.Errorf("failed to update reorg LendingTrade. Err: %v . Trade: %s", err.Error(), lendingstate.ToJSON(trade))
			}
		}
	}

	// remove repay/topup history
	db.DeleteItemByTxHash(txhash, &lendingstate.LendingItem{Status: lendingstate.Repay})
	db.DeleteItemByTxHash(txhash, &lendingstate.LendingItem{Status: lendingstate.TopUp})

	if err := db.CommitLendingBulk(); err != nil {
		return fmt.Errorf("failed to RollbackLendingData. %v", err)
	}
	return nil
}

func (l *Lending) ProcessLiquidationData(chain consensus.ChainContext, time *big.Int, statedb *state.StateDB, tradingState *tradingstate.TradingStateDB, lendingState *lendingstate.LendingStateDB) (finalizedTrades map[common.Hash]*lendingstate.LendingTrade, err error) {
	allPairs, err := lendingstate.GetAllLendingPairs(statedb)
	if err != nil {
		log.Debug("Not found all trading pairs", "error", err)
		return map[common.Hash]*lendingstate.LendingTrade{}, nil
	}
	allLendingBooks, err := lendingstate.GetAllLendingBooks(statedb)
	if err != nil {
		log.Debug("Not found all lending books", "error", err)
		return map[common.Hash]*lendingstate.LendingTrade{}, nil
	}

	finalizedTrades = map[common.Hash]*lendingstate.LendingTrade{}

	// liquidate trades by time
	for lendingBook, _ := range allLendingBooks {
		lowestTime, tradingIds := lendingState.GetLowestLiquidationTime(lendingBook, time)
		log.Debug("ProcessLiquidationData time", "tradeIds", len(tradingIds))
		for lowestTime.Sign() > 0 && lowestTime.Cmp(time) < 0 {
			for _, tradingId := range tradingIds {
				log.Debug("ProcessRepay", "lowestTime", lowestTime, "time", time, "lendingBook", lendingBook.Hex(), "tradingId", tradingId.Hex())
				trade, err := l.ProcessRepay(time.Uint64(), lendingState, statedb, tradingState, lendingBook, tradingId.Big().Uint64())
				if err != nil {
					log.Error("Fail when process payment ", "time", time, "lendingBook", lendingBook.Hex(), "tradingId", tradingId, "error", err)
					return map[common.Hash]*lendingstate.LendingTrade{}, err
				}
				if trade != nil && trade.Hash != (common.Hash{}) {
					finalizedTrades[trade.Hash] = trade
				}
			}
			lowestTime, tradingIds = lendingState.GetLowestLiquidationTime(lendingBook, time)
		}
	}

	// liquidate trades by collateralPrice
	for _, lendingPair := range allPairs {
		orderbook := tradingstate.GetTradingOrderBookHash(lendingPair.CollateralToken, lendingPair.LendingToken)
		_, liquidationPrice, err := l.GetCollateralPrices(chain, statedb, tradingState, lendingPair.CollateralToken, lendingPair.LendingToken)
		if err != nil {
			log.Error("Fail when get all trading pairs", "error", err)
			return map[common.Hash]*lendingstate.LendingTrade{}, err
		}
		highestPrice, liquidationData := tradingState.GetHighestLiquidationPriceData(orderbook, liquidationPrice)
		for highestPrice.Sign() > 0 && liquidationPrice.Cmp(highestPrice) < 0 {
			for lendingBook, tradingIds := range liquidationData {
				for _, tradingIdHash := range tradingIds {
					trade := lendingState.GetLendingTrade(lendingBook, tradingIdHash)
					if trade.AutoTopUp {
						if newTrade, err := l.AutoTopUp(statedb, tradingState, lendingState, lendingBook, tradingIdHash, liquidationPrice); err == nil {
							// if this action complete successfully, do not liquidate this trade in this epoch
							log.Debug("AutoTopUp", "borrower", trade.Borrower.Hex(), "collateral", newTrade.CollateralToken.Hex(), "newLockedAmount", newTrade.CollateralLockedAmount)
							finalizedTrades[newTrade.Hash] = newTrade
							continue
						}
					}
					log.Debug("LiquidationTrade", "highestPrice", highestPrice, "lendingBook", lendingBook.Hex(), "tradingIdHash", tradingIdHash.Hex())
					newTrade, err := l.LiquidationTrade(lendingState, statedb, tradingState, lendingBook, tradingIdHash.Big().Uint64())
					if err != nil {
						log.Error("Fail when remove liquidation newTrade", "time", time, "lendingBook", lendingBook.Hex(), "tradingIdHash", tradingIdHash.Hex(), "error", err)
						return map[common.Hash]*lendingstate.LendingTrade{}, err
					}
					if newTrade != nil && newTrade.Hash != (common.Hash{}) {
						newTrade.Status = lendingstate.TradeStatusLiquidated
						finalizedTrades[newTrade.Hash] = newTrade
					}
				}
			}
			highestPrice, liquidationData = tradingState.GetHighestLiquidationPriceData(orderbook, liquidationPrice)
		}
	}

	log.Debug("ProcessLiquidationData", "len", len(finalizedTrades))
	return finalizedTrades, nil
}
