package tomox

import (
	"github.com/rupaya-project/rupx/contracts/tomox/contract"
	"github.com/rupaya-project/rupx/log"
	"math/big"
	"strings"

	"github.com/rupaya-project/rupx"
	"github.com/rupaya-project/rupx/accounts/abi"
	"github.com/rupaya-project/rupx/accounts/abi/bind/backends"
	"github.com/rupaya-project/rupx/common"
	"github.com/rupaya-project/rupx/consensus"
	"github.com/rupaya-project/rupx/core/state"
)


// GetTokenAbi return token abi
func GetTokenAbi() (*abi.ABI, error) {
	contractABI, err := abi.JSON(strings.NewReader(contract.TRC21ABI))
	if err != nil {
		return nil, err
	}
	return &contractABI, nil
}

// RunContract run smart contract
func RunContract(chain consensus.ChainContext, statedb *state.StateDB, contractAddr common.Address, abi *abi.ABI, method string, args ...interface{}) (interface{}, error) {
	input, err := abi.Pack(method)
	if err != nil {
		return nil, err
	}
	backend := (*backends.SimulatedBackend)(nil)
	fakeCaller := common.HexToAddress("0x0000000000000000000000000000000000000001")
	msg := tomochain.CallMsg{To: &contractAddr, Data: input, From: fakeCaller}
	result, err := backend.CallContractWithState(msg, chain, statedb)
	if err != nil {
		return nil, err
	}
	var unpackResult interface{}
	err = abi.Unpack(&unpackResult, method, result)
	if err != nil {
		return nil, err
	}
	return unpackResult, nil
}

func (tomox *TomoX) GetTokenDecimal(chain consensus.ChainContext, statedb *state.StateDB, tokenAddr common.Address) (*big.Int, error) {
	if tokenDecimal, ok := tomox.tokenDecimalCache.Get(tokenAddr); ok {
		return tokenDecimal.(*big.Int), nil
	}
	if tokenAddr.String() == common.TomoNativeAddress {
		tomox.tokenDecimalCache.Add(tokenAddr, common.BasePrice)
		return common.BasePrice, nil
	}
	var decimals uint8
	defer func() {
		log.Debug("GetTokenDecimal from ", "relayerSMC", common.RelayerRegistrationSMC, "tokenAddr", tokenAddr.Hex(), "decimals", decimals)
	}()
	contractABI, err := GetTokenAbi()
	if err != nil {
		return nil, err
	}
	stateCopy := statedb.Copy()
	result, err := RunContract(chain, stateCopy, tokenAddr, contractABI, "decimals")
	if err != nil {
		return nil, err
	}
	decimals = result.(uint8)

	tokenDecimal := new(big.Int).SetUint64(0).Exp(big.NewInt(10), big.NewInt(int64(decimals)), nil)
	tomox.tokenDecimalCache.Add(tokenAddr, tokenDecimal)
	return tokenDecimal, nil
}
