package rrc21issuer

import (
	"github.com/rupaya-project/rupx/accounts/abi/bind"
	"github.com/rupaya-project/rupx/common"
	"github.com/rupaya-project/rupx/contracts/rrc21issuer/contract"
	"math/big"
)

type RRC21Issuer struct {
	*contract.RRC21IssuerSession
	contractBackend bind.ContractBackend
}

func NewRRC21Issuer(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*RRC21Issuer, error) {
	contractObject, err := contract.NewRRC21Issuer(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &RRC21Issuer{
		&contract.RRC21IssuerSession{
			Contract:     contractObject,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployRRC21Issuer(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend, minApply *big.Int) (common.Address, *RRC21Issuer, error) {
	contractAddr, _, _, err := contract.DeployRRC21Issuer(transactOpts, contractBackend, minApply)
	if err != nil {
		return contractAddr, nil, err
	}
	contractObject, err := NewRRC21Issuer(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, contractObject, nil
}
