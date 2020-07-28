package rupex

import (
	"errors"
	"fmt"
	"math/big"
	"strconv"
	"time"

	"github.com/rupaya-project/rupx/consensus"
	"github.com/rupaya-project/rupx/core/types"
	"github.com/rupaya-project/rupx/p2p"
	"github.com/rupaya-project/rupx/rupex/tradingstate"
	rupexDAO "github.com/rupaya-project/rupx/rupexDAO"
	"gopkg.in/karalabe/cookiejar.v2/collections/prque"

	lru "github.com/hashicorp/golang-lru"
	"github.com/rupaya-project/rupx/common"
	"github.com/rupaya-project/rupx/core/state"
	"github.com/rupaya-project/rupx/log"
	"github.com/rupaya-project/rupx/rpc"
	"golang.org/x/sync/syncmap"
)

const (
	ProtocolName       = "rupex"
	ProtocolVersion    = uint64(1)
	ProtocolVersionStr = "1.0"
	overflowIdx        // Indicator of message queue overflow
	defaultCacheLimit  = 1024
	MaximumTxMatchSize = 1000
)

var (
	ErrNonceTooHigh = errors.New("nonce too high")
	ErrNonceTooLow  = errors.New("nonce too low")
)

type Config struct {
	DataDir        string `toml:",omitempty"`
	DBEngine       string `toml:",omitempty"`
	DBName         string `toml:",omitempty"`
	ConnectionUrl  string `toml:",omitempty"`
	ReplicaSetName string `toml:",omitempty"`
}

// DefaultConfig represents (shocker!) the default configuration.
var DefaultConfig = Config{
	DataDir: "",
}

type RupeX struct {
	// Order related
	db         rupexDAO.RupeXDAO
	mongodb    rupexDAO.RupeXDAO
	Triegc     *prque.Prque          // Priority queue mapping block numbers to tries to gc
	StateCache tradingstate.Database // State database to reuse between imports (contains state cache)    *rupex_state.TradingStateDB

	orderNonce map[common.Address]*big.Int

	sdkNode           bool
	settings          syncmap.Map // holds configuration settings that can be dynamically changed
	tokenDecimalCache *lru.Cache
	orderCache        *lru.Cache
}

func (rupex *RupeX) Protocols() []p2p.Protocol {
	return []p2p.Protocol{}
}

func (rupex *RupeX) Start(server *p2p.Server) error {
	return nil
}

func (rupex *RupeX) SaveData() {
}
func (rupex *RupeX) Stop() error {
	return nil
}

func NewLDBEngine(cfg *Config) *rupexDAO.BatchDatabase {
	datadir := cfg.DataDir
	batchDB := rupexDAO.NewBatchDatabaseWithEncode(datadir, 0)
	return batchDB
}

func NewMongoDBEngine(cfg *Config) *rupexDAO.MongoDatabase {
	mongoDB, err := rupexDAO.NewMongoDatabase(nil, cfg.DBName, cfg.ConnectionUrl, cfg.ReplicaSetName, 0)

	if err != nil {
		log.Crit("Failed to init mongodb engine", "err", err)
	}

	return mongoDB
}

func New(cfg *Config) *RupeX {
	tokenDecimalCache, _ := lru.New(defaultCacheLimit)
	orderCache, _ := lru.New(tradingstate.OrderCacheLimit)
	rupayaX := &RupeX{
		orderNonce:        make(map[common.Address]*big.Int),
		Triegc:            prque.New(),
		tokenDecimalCache: tokenDecimalCache,
		orderCache:        orderCache,
	}

	// default DBEngine: levelDB
	rupayaX.db = NewLDBEngine(cfg)
	rupayaX.sdkNode = false

	if cfg.DBEngine == "mongodb" { // this is an add-on DBEngine for SDK nodes
		rupayaX.mongodb = NewMongoDBEngine(cfg)
		rupayaX.sdkNode = true
	}

	rupayaX.StateCache = tradingstate.NewDatabase(rupayaX.db)
	rupayaX.settings.Store(overflowIdx, false)

	return rupayaX
}

// Overflow returns an indication if the message queue is full.
func (rupex *RupeX) Overflow() bool {
	val, _ := rupex.settings.Load(overflowIdx)
	return val.(bool)
}

func (rupex *RupeX) IsSDKNode() bool {
	return rupex.sdkNode
}

func (rupex *RupeX) GetLevelDB() rupexDAO.RupeXDAO {
	return rupex.db
}

func (rupex *RupeX) GetMongoDB() rupexDAO.RupeXDAO {
	return rupex.mongodb
}

// APIs returns the RPC descriptors the RupeX implementation offers
func (rupex *RupeX) APIs() []rpc.API {
	return []rpc.API{
		{
			Namespace: ProtocolName,
			Version:   ProtocolVersionStr,
			Service:   NewPublicRupeXAPI(rupex),
			Public:    true,
		},
	}
}

// Version returns the RupeX sub-protocols version number.
func (rupex *RupeX) Version() uint64 {
	return ProtocolVersion
}

func (rupex *RupeX) ProcessOrderPending(coinbase common.Address, chain consensus.ChainContext, pending map[common.Address]types.OrderTransactions, statedb *state.StateDB, rupayaXstatedb *tradingstate.TradingStateDB) ([]tradingstate.TxDataMatch, map[common.Hash]tradingstate.MatchingResult) {
	txMatches := []tradingstate.TxDataMatch{}
	matchingResults := map[common.Hash]tradingstate.MatchingResult{}

	txs := types.NewOrderTransactionByNonce(types.OrderTxSigner{}, pending)
	numberTx := 0
	for {
		tx := txs.Peek()
		if tx == nil {
			break
		}
		if numberTx > MaximumTxMatchSize {
			break
		}
		numberTx++
		log.Debug("ProcessOrderPending start", "len", len(pending))
		log.Debug("Get pending orders to process", "address", tx.UserAddress(), "nonce", tx.Nonce())
		V, R, S := tx.Signature()

		bigstr := V.String()
		n, e := strconv.ParseInt(bigstr, 10, 8)
		if e != nil {
			continue
		}

		order := &tradingstate.OrderItem{
			Nonce:           big.NewInt(int64(tx.Nonce())),
			Quantity:        tx.Quantity(),
			Price:           tx.Price(),
			ExchangeAddress: tx.ExchangeAddress(),
			UserAddress:     tx.UserAddress(),
			BaseToken:       tx.BaseToken(),
			QuoteToken:      tx.QuoteToken(),
			Status:          tx.Status(),
			Side:            tx.Side(),
			Type:            tx.Type(),
			Hash:            tx.OrderHash(),
			OrderID:         tx.OrderID(),
			Signature: &tradingstate.Signature{
				V: byte(n),
				R: common.BigToHash(R),
				S: common.BigToHash(S),
			},
		}
		cancel := false
		if order.Status == tradingstate.OrderStatusCancelled {
			cancel = true
		}

		log.Info("Process order pending", "orderPending", order, "BaseToken", order.BaseToken.Hex(), "QuoteToken", order.QuoteToken)
		originalOrder := &tradingstate.OrderItem{}
		*originalOrder = *order
		originalOrder.Quantity = tradingstate.CloneBigInt(order.Quantity)

		if cancel {
			order.Status = tradingstate.OrderStatusCancelled
		}

		newTrades, newRejectedOrders, err := rupex.CommitOrder(coinbase, chain, statedb, rupayaXstatedb, tradingstate.GetTradingOrderBookHash(order.BaseToken, order.QuoteToken), order)

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
		originalOrder.OrderID = order.OrderID
		originalOrder.ExtraData = order.ExtraData
		originalOrderValue, err := tradingstate.EncodeBytesItem(originalOrder)
		if err != nil {
			log.Error("Can't encode", "order", originalOrder, "err", err)
			continue
		}
		txMatch := tradingstate.TxDataMatch{
			Order: originalOrderValue,
		}
		txMatches = append(txMatches, txMatch)
		matchingResults[tradingstate.GetMatchingResultCacheKey(order)] = tradingstate.MatchingResult{
			Trades:  newTrades,
			Rejects: newRejectedOrders,
		}
	}
	return txMatches, matchingResults
}

// there are 3 tasks need to complete to update data in SDK nodes after matching
// 1. txMatchData.Order: order has been processed. This order should be put to `orders` collection with status sdktypes.OrderStatusOpen
// 2. txMatchData.Trades: includes information of matched orders.
// 		a. PutObject them to `trades` collection
// 		b. Update status of regrading orders to sdktypes.OrderStatusFilled
func (rupex *RupeX) SyncDataToSDKNode(takerOrderInTx *tradingstate.OrderItem, txHash common.Hash, txMatchTime time.Time, statedb *state.StateDB, trades []map[string]string, rejectedOrders []*tradingstate.OrderItem, dirtyOrderCount *uint64) error {
	var (
		// originTakerOrder: order get from db, nil if it doesn't exist
		// takerOrderInTx: order decoded from txdata
		// updatedTakerOrder: order with new status, filledAmount, CreatedAt, UpdatedAt. This will be inserted to db
		originTakerOrder, updatedTakerOrder *tradingstate.OrderItem
		makerDirtyHashes                    []string
		makerDirtyFilledAmount              map[string]*big.Int
		err                                 error
	)
	db := rupex.GetMongoDB()
	db.InitBulk()
	if takerOrderInTx.Status == tradingstate.OrderStatusCancelled && len(rejectedOrders) > 0 {
		// cancel order is rejected -> nothing change
		log.Debug("Cancel order is rejected", "order", tradingstate.ToJSON(takerOrderInTx))
		return nil
	}
	// 1. put processed takerOrderInTx to db
	lastState := tradingstate.OrderHistoryItem{}
	val, err := db.GetObject(takerOrderInTx.Hash, &tradingstate.OrderItem{})
	if err == nil && val != nil {
		originTakerOrder = val.(*tradingstate.OrderItem)
		lastState = tradingstate.OrderHistoryItem{
			TxHash:       originTakerOrder.TxHash,
			FilledAmount: tradingstate.CloneBigInt(originTakerOrder.FilledAmount),
			Status:       originTakerOrder.Status,
			UpdatedAt:    originTakerOrder.UpdatedAt,
		}
	}
	if originTakerOrder != nil {
		updatedTakerOrder = originTakerOrder
	} else {
		updatedTakerOrder = takerOrderInTx
		updatedTakerOrder.FilledAmount = new(big.Int)
	}

	if takerOrderInTx.Status != tradingstate.OrderStatusCancelled {
		updatedTakerOrder.Status = tradingstate.OrderStatusOpen
	} else {
		updatedTakerOrder.Status = tradingstate.OrderStatusCancelled
		updatedTakerOrder.ExtraData = takerOrderInTx.ExtraData
	}
	updatedTakerOrder.TxHash = txHash
	if updatedTakerOrder.CreatedAt.IsZero() {
		updatedTakerOrder.CreatedAt = txMatchTime
	}
	if txMatchTime.Before(updatedTakerOrder.UpdatedAt) || (txMatchTime.Equal(updatedTakerOrder.UpdatedAt) && *dirtyOrderCount == 0) {
		log.Debug("Ignore old orders/trades taker", "txHash", txHash.Hex(), "txTime", txMatchTime.UnixNano(), "updatedAt", updatedTakerOrder.UpdatedAt.UnixNano())
		return nil
	}
	*dirtyOrderCount++

	rupex.UpdateOrderCache(updatedTakerOrder.BaseToken, updatedTakerOrder.QuoteToken, updatedTakerOrder.Hash, txHash, lastState)
	updatedTakerOrder.UpdatedAt = txMatchTime

	// 2. put trades to db and update status to FILLED
	log.Debug("Got trades", "number", len(trades), "txhash", txHash.Hex())
	makerDirtyFilledAmount = make(map[string]*big.Int)
	for _, trade := range trades {
		// 2.a. put to trades
		if trade == nil {
			continue
		}
		tradeRecord := &tradingstate.Trade{}
		quantity := tradingstate.ToBigInt(trade[tradingstate.TradeQuantity])
		price := tradingstate.ToBigInt(trade[tradingstate.TradePrice])
		if price.Cmp(big.NewInt(0)) <= 0 || quantity.Cmp(big.NewInt(0)) <= 0 {
			return fmt.Errorf("trade misses important information. tradedPrice %v, tradedQuantity %v", price, quantity)
		}
		tradeRecord.Amount = quantity
		tradeRecord.PricePoint = price
		tradeRecord.BaseToken = updatedTakerOrder.BaseToken
		tradeRecord.QuoteToken = updatedTakerOrder.QuoteToken
		tradeRecord.Status = tradingstate.TradeStatusSuccess
		tradeRecord.Taker = updatedTakerOrder.UserAddress
		tradeRecord.Maker = common.HexToAddress(trade[tradingstate.TradeMaker])
		tradeRecord.TakerOrderHash = updatedTakerOrder.Hash
		tradeRecord.MakerOrderHash = common.HexToHash(trade[tradingstate.TradeMakerOrderHash])
		tradeRecord.TxHash = txHash
		tradeRecord.TakerOrderSide = updatedTakerOrder.Side
		tradeRecord.TakerExchange = updatedTakerOrder.ExchangeAddress
		tradeRecord.MakerExchange = common.HexToAddress(trade[tradingstate.TradeMakerExchange])

		tradeRecord.MakeFee, _ = new(big.Int).SetString(trade[tradingstate.MakerFee], 10)
		tradeRecord.TakeFee, _ = new(big.Int).SetString(trade[tradingstate.TakerFee], 10)

		// set makerOrderType, takerOrderType
		tradeRecord.MakerOrderType = trade[tradingstate.MakerOrderType]
		tradeRecord.TakerOrderType = updatedTakerOrder.Type

		if tradeRecord.CreatedAt.IsZero() {
			tradeRecord.CreatedAt = txMatchTime
		}
		tradeRecord.UpdatedAt = txMatchTime
		tradeRecord.Hash = tradeRecord.ComputeHash()

		log.Debug("TRADE history", "amount", tradeRecord.Amount, "pricepoint", tradeRecord.PricePoint,
			"taker", tradeRecord.Taker.Hex(), "maker", tradeRecord.Maker.Hex(), "takerOrder", tradeRecord.TakerOrderHash.Hex(), "makerOrder", tradeRecord.MakerOrderHash.Hex(),
			"takerFee", tradeRecord.TakeFee, "makerFee", tradeRecord.MakeFee)
		if err := db.PutObject(tradeRecord.Hash, tradeRecord); err != nil {
			return fmt.Errorf("SDKNode: failed to store tradeRecord %s", err.Error())
		}

		// 2.b. update status and filledAmount
		filledAmount := quantity
		// maker dirty order
		makerFilledAmount := big.NewInt(0)
		if amount, ok := makerDirtyFilledAmount[trade[tradingstate.TradeMakerOrderHash]]; ok {
			makerFilledAmount = tradingstate.CloneBigInt(amount)
		}
		makerFilledAmount = new(big.Int).Add(makerFilledAmount, filledAmount)
		makerDirtyFilledAmount[trade[tradingstate.TradeMakerOrderHash]] = makerFilledAmount
		makerDirtyHashes = append(makerDirtyHashes, trade[tradingstate.TradeMakerOrderHash])

		//updatedTakerOrder = rupex.updateMatchedOrder(updatedTakerOrder, filledAmount, txMatchTime, txHash)
		//  update filledAmount, status of takerOrder
		updatedTakerOrder.FilledAmount = new(big.Int).Add(updatedTakerOrder.FilledAmount, filledAmount)
		if updatedTakerOrder.FilledAmount.Cmp(updatedTakerOrder.Quantity) < 0 && updatedTakerOrder.Type == tradingstate.Limit {
			updatedTakerOrder.Status = tradingstate.OrderStatusPartialFilled
		} else {
			updatedTakerOrder.Status = tradingstate.OrderStatusFilled
		}
	}

	// for Market orders
	// filledAmount > 0 : FILLED
	// otherwise: REJECTED
	if updatedTakerOrder.Type == tradingstate.Market {
		if updatedTakerOrder.FilledAmount.Sign() > 0 {
			updatedTakerOrder.Status = tradingstate.OrderStatusFilled
		} else {
			updatedTakerOrder.Status = tradingstate.OrderStatusRejected
		}
	}
	log.Debug("PutObject processed takerOrder",
		"userAddr", updatedTakerOrder.UserAddress.Hex(), "side", updatedTakerOrder.Side,
		"price", updatedTakerOrder.Price, "quantity", updatedTakerOrder.Quantity, "filledAmount", updatedTakerOrder.FilledAmount, "status", updatedTakerOrder.Status,
		"hash", updatedTakerOrder.Hash.Hex(), "txHash", updatedTakerOrder.TxHash.Hex())
	if err := db.PutObject(updatedTakerOrder.Hash, updatedTakerOrder); err != nil {
		return fmt.Errorf("SDKNode: failed to put processed takerOrder. Hash: %s Error: %s", updatedTakerOrder.Hash.Hex(), err.Error())
	}
	items := db.GetListItemByHashes(makerDirtyHashes, &tradingstate.OrderItem{})
	if items != nil {
		makerOrders := items.([]*tradingstate.OrderItem)
		log.Debug("Maker dirty orders", "len", len(makerOrders), "txhash", txHash.Hex())
		for _, o := range makerOrders {
			if txMatchTime.Before(o.UpdatedAt) {
				log.Debug("Ignore old orders/trades maker", "txHash", txHash.Hex(), "txTime", txMatchTime.UnixNano(), "updatedAt", updatedTakerOrder.UpdatedAt.UnixNano())
				continue
			}
			lastState = tradingstate.OrderHistoryItem{
				TxHash:       o.TxHash,
				FilledAmount: tradingstate.CloneBigInt(o.FilledAmount),
				Status:       o.Status,
				UpdatedAt:    o.UpdatedAt,
			}
			rupex.UpdateOrderCache(o.BaseToken, o.QuoteToken, o.Hash, txHash, lastState)
			o.TxHash = txHash
			o.UpdatedAt = txMatchTime
			o.FilledAmount = new(big.Int).Add(o.FilledAmount, makerDirtyFilledAmount[o.Hash.Hex()])
			if o.FilledAmount.Cmp(o.Quantity) < 0 {
				o.Status = tradingstate.OrderStatusPartialFilled
			} else {
				o.Status = tradingstate.OrderStatusFilled
			}
			log.Debug("PutObject processed makerOrder",
				"userAddr", o.UserAddress.Hex(), "side", o.Side,
				"price", o.Price, "quantity", o.Quantity, "filledAmount", o.FilledAmount, "status", o.Status,
				"hash", o.Hash.Hex(), "txHash", o.TxHash.Hex())
			if err := db.PutObject(o.Hash, o); err != nil {
				return fmt.Errorf("SDKNode: failed to put processed makerOrder. Hash: %s Error: %s", o.Hash.Hex(), err.Error())
			}
		}
	}

	// 3. put rejected orders to db and update status REJECTED
	log.Debug("Got rejected orders", "number", len(rejectedOrders), "rejectedOrders", rejectedOrders)

	if len(rejectedOrders) > 0 {
		var rejectedHashes []string
		// updateRejectedOrders
		for _, rejectedOrder := range rejectedOrders {
			rejectedHashes = append(rejectedHashes, rejectedOrder.Hash.Hex())
			if updatedTakerOrder.Hash == rejectedOrder.Hash && !txMatchTime.Before(updatedTakerOrder.UpdatedAt) {
				// cache order history for handling reorg
				orderHistoryRecord := tradingstate.OrderHistoryItem{
					TxHash:       updatedTakerOrder.TxHash,
					FilledAmount: tradingstate.CloneBigInt(updatedTakerOrder.FilledAmount),
					Status:       updatedTakerOrder.Status,
					UpdatedAt:    updatedTakerOrder.UpdatedAt,
				}
				rupex.UpdateOrderCache(updatedTakerOrder.BaseToken, updatedTakerOrder.QuoteToken, updatedTakerOrder.Hash, txHash, orderHistoryRecord)
				// if whole order is rejected, status = REJECTED
				// otherwise, status = FILLED
				if updatedTakerOrder.FilledAmount.Sign() > 0 {
					updatedTakerOrder.Status = tradingstate.OrderStatusFilled
				} else {
					updatedTakerOrder.Status = tradingstate.OrderStatusRejected
				}
				updatedTakerOrder.TxHash = txHash
				updatedTakerOrder.UpdatedAt = txMatchTime
				if err := db.PutObject(updatedTakerOrder.Hash, updatedTakerOrder); err != nil {
					return fmt.Errorf("SDKNode: failed to reject takerOrder. Hash: %s Error: %s", updatedTakerOrder.Hash.Hex(), err.Error())
				}
			}
		}
		items := db.GetListItemByHashes(rejectedHashes, &tradingstate.OrderItem{})
		if items != nil {
			dirtyRejectedOrders := items.([]*tradingstate.OrderItem)
			for _, order := range dirtyRejectedOrders {
				if txMatchTime.Before(order.UpdatedAt) {
					log.Debug("Ignore old orders/trades reject", "txHash", txHash.Hex(), "txTime", txMatchTime.UnixNano(), "updatedAt", updatedTakerOrder.UpdatedAt.UnixNano())
					continue
				}
				// cache order history for handling reorg
				orderHistoryRecord := tradingstate.OrderHistoryItem{
					TxHash:       order.TxHash,
					FilledAmount: tradingstate.CloneBigInt(order.FilledAmount),
					Status:       order.Status,
					UpdatedAt:    order.UpdatedAt,
				}
				rupex.UpdateOrderCache(order.BaseToken, order.QuoteToken, order.Hash, txHash, orderHistoryRecord)
				dirtyFilledAmount, ok := makerDirtyFilledAmount[order.Hash.Hex()]
				if ok && dirtyFilledAmount != nil {
					order.FilledAmount = new(big.Int).Add(order.FilledAmount, dirtyFilledAmount)
				}
				// if whole order is rejected, status = REJECTED
				// otherwise, status = FILLED
				if order.FilledAmount.Sign() > 0 {
					order.Status = tradingstate.OrderStatusFilled
				} else {
					order.Status = tradingstate.OrderStatusRejected
				}
				order.TxHash = txHash
				order.UpdatedAt = txMatchTime
				if err = db.PutObject(order.Hash, order); err != nil {
					return fmt.Errorf("SDKNode: failed to update rejectedOder to sdkNode %s", err.Error())
				}
			}
		}
	}

	if err := db.CommitBulk(); err != nil {
		return fmt.Errorf("SDKNode fail to commit bulk update orders, trades at txhash %s . Error: %s", txHash.Hex(), err.Error())
	}
	return nil
}

func (rupex *RupeX) GetTradingState(block *types.Block, author common.Address) (*tradingstate.TradingStateDB, error) {
	root, err := rupex.GetTradingStateRoot(block, author)
	if err != nil {
		return nil, err
	}
	if rupex.StateCache == nil {
		return nil, errors.New("Not initialized rupex")
	}
	return tradingstate.New(root, rupex.StateCache)
}

func (rupex *RupeX) GetStateCache() tradingstate.Database {
	return rupex.StateCache
}
func (rupex *RupeX) HasTradingState(block *types.Block, author common.Address) bool {
	root, err := rupex.GetTradingStateRoot(block, author)
	if err != nil {
		return false
	}
	_, err = rupex.StateCache.OpenTrie(root)
	if err != nil {
		return false
	}
	return true
}
func (rupex *RupeX) GetTriegc() *prque.Prque {
	return rupex.Triegc
}

func (rupex *RupeX) GetTradingStateRoot(block *types.Block, author common.Address) (common.Hash, error) {
	for _, tx := range block.Transactions() {
		from := *(tx.From())
		if tx.To() != nil && tx.To().Hex() == common.TradingStateAddr && from.String() == author.String() {
			if len(tx.Data()) >= 32 {
				return common.BytesToHash(tx.Data()[:32]), nil
			}
		}
	}
	return tradingstate.EmptyRoot, nil
}

func (rupex *RupeX) UpdateOrderCache(baseToken, quoteToken common.Address, orderHash common.Hash, txhash common.Hash, lastState tradingstate.OrderHistoryItem) {
	var orderCacheAtTxHash map[common.Hash]tradingstate.OrderHistoryItem
	c, ok := rupex.orderCache.Get(txhash)
	if !ok || c == nil {
		orderCacheAtTxHash = make(map[common.Hash]tradingstate.OrderHistoryItem)
	} else {
		orderCacheAtTxHash = c.(map[common.Hash]tradingstate.OrderHistoryItem)
	}
	orderKey := tradingstate.GetOrderHistoryKey(baseToken, quoteToken, orderHash)
	_, ok = orderCacheAtTxHash[orderKey]
	if !ok {
		orderCacheAtTxHash[orderKey] = lastState
	}
	rupex.orderCache.Add(txhash, orderCacheAtTxHash)
}

func (rupex *RupeX) RollbackReorgTxMatch(txhash common.Hash) error {
	db := rupex.GetMongoDB()
	db.InitBulk()

	items := db.GetListItemByTxHash(txhash, &tradingstate.OrderItem{})
	if items != nil {
		for _, order := range items.([]*tradingstate.OrderItem) {
			c, ok := rupex.orderCache.Get(txhash)
			log.Debug("Rupex reorg: rollback order", "txhash", txhash.Hex(), "order", tradingstate.ToJSON(order), "orderHistoryItem", c)
			if !ok {
				log.Debug("Rupex reorg: remove order due to no orderCache", "order", tradingstate.ToJSON(order))
				if err := db.DeleteObject(order.Hash, &tradingstate.OrderItem{}); err != nil {
					log.Crit("SDKNode: failed to remove reorg order", "err", err.Error(), "order", tradingstate.ToJSON(order))
				}
				continue
			}
			orderCacheAtTxHash := c.(map[common.Hash]tradingstate.OrderHistoryItem)
			orderHistoryItem, _ := orderCacheAtTxHash[tradingstate.GetOrderHistoryKey(order.BaseToken, order.QuoteToken, order.Hash)]
			if (orderHistoryItem == tradingstate.OrderHistoryItem{}) {
				log.Debug("Rupex reorg: remove order due to empty orderHistory", "order", tradingstate.ToJSON(order))
				if err := db.DeleteObject(order.Hash, &tradingstate.OrderItem{}); err != nil {
					log.Crit("SDKNode: failed to remove reorg order", "err", err.Error(), "order", tradingstate.ToJSON(order))
				}
				continue
			}
			order.TxHash = orderHistoryItem.TxHash
			order.Status = orderHistoryItem.Status
			order.FilledAmount = tradingstate.CloneBigInt(orderHistoryItem.FilledAmount)
			order.UpdatedAt = orderHistoryItem.UpdatedAt
			log.Debug("Rupex reorg: update order to the last orderHistoryItem", "order", tradingstate.ToJSON(order), "orderHistoryItem", orderHistoryItem)
			if err := db.PutObject(order.Hash, order); err != nil {
				log.Crit("SDKNode: failed to update reorg order", "err", err.Error(), "order", tradingstate.ToJSON(order))
			}
		}
	}
	log.Debug("Rupex reorg: DeleteTradeByTxHash", "txhash", txhash.Hex())
	db.DeleteItemByTxHash(txhash, &tradingstate.Trade{})
	if err := db.CommitBulk(); err != nil {
		return fmt.Errorf("failed to RollbackTradingData. %v", err)
	}
	return nil
}
