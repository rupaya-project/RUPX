package rupex

import (
	"github.com/rupaya-project/rupx/accounts/abi/bind"
	"github.com/rupaya-project/rupx/common"
	"github.com/rupaya-project/rupx/contracts/rupex/contract"
)

type RUPEXListing struct {
	*contract.RUPEXListingSession
	contractBackend bind.ContractBackend
}

func NewMyRUPEXListing(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*RUPEXListing, error) {
	smartContract, err := contract.NewRUPEXListing(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &RUPEXListing{
		&contract.RUPEXListingSession{
			Contract:     smartContract,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployRUPEXListing(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend) (common.Address, *RUPEXListing, error) {
	contractAddr, _, _, err := contract.DeployRUPEXListing(transactOpts, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}
	smartContract, err := NewMyRUPEXListing(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, smartContract, nil
}
