package main

import (
	"context"
	"fmt"
	"github.com/rupaya-project/rupx/accounts/abi/bind"
	"github.com/rupaya-project/rupx/common"
	"github.com/rupaya-project/rupx/contracts/rrc21issuer"
	"github.com/rupaya-project/rupx/contracts/rrc21issuer/simulation"
	"github.com/rupaya-project/rupx/ethclient"
	"log"
	"math/big"
	"time"
)

func main() {
	fmt.Println("========================")
	fmt.Println("mainAddr", simulation.MainAddr.Hex())
	fmt.Println("airdropAddr", simulation.AirdropAddr.Hex())
	fmt.Println("receiverAddr", simulation.ReceiverAddr.Hex())
	fmt.Println("========================")
	client, err := ethclient.Dial(simulation.RpcEndpoint)
	if err != nil {
		fmt.Println(err, client)
	}
	nonce, _ := client.NonceAt(context.Background(), simulation.MainAddr, nil)

	// init rrc21 issuer
	auth := bind.NewKeyedTransactor(simulation.MainKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(4000000) // in units
	auth.GasPrice = big.NewInt(210000000000000)
	rrc21IssuerAddr, rrc21Issuer, err := rrc21issuer.DeployRRC21Issuer(auth, client, simulation.MinApply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("main address", simulation.MainAddr.Hex(), "nonce", nonce)
	fmt.Println("===> rrc21 issuer address", rrc21IssuerAddr.Hex())

	auth.Nonce = big.NewInt(int64(nonce + 1))

	// init rrc20
	rrc21TokenAddr, rrc21Token, err := rrc21issuer.DeployRRC21(auth, client, "TEST", "RUPX", 18, simulation.Cap, simulation.Fee)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("===>  rrc21 token address", rrc21TokenAddr.Hex(), "cap", simulation.Cap)

	fmt.Println("wait 10s to execute init smart contract")
	time.Sleep(10 * time.Second)

	rrc21Issuer.TransactOpts.Nonce = big.NewInt(int64(nonce + 2))
	rrc21Issuer.TransactOpts.Value = simulation.MinApply
	rrc21Issuer.TransactOpts.GasPrice = big.NewInt(common.DefaultMinGasPrice)
	rrc21Token.TransactOpts.GasPrice = big.NewInt(common.DefaultMinGasPrice)
	rrc21Token.TransactOpts.GasLimit = uint64(4000000)
	auth.GasPrice = big.NewInt(common.DefaultMinGasPrice)
	// get balance init rrc21 smart contract
	balance, err := rrc21Token.BalanceOf(simulation.MainAddr)
	if err != nil || balance.Cmp(simulation.Cap) != 0 {
		log.Fatal(err, "\tget\t", balance, "\twant\t", simulation.Cap)
	}
	fmt.Println("balance", balance, "mainAddr", simulation.MainAddr.Hex())

	// add rrc20 list token rrc21 issuer
	_, err = rrc21Issuer.Apply(rrc21TokenAddr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("wait 10s to add token to list issuer")
	time.Sleep(10 * time.Second)

	//check rrc21 SMC balance
	balance, err = client.BalanceAt(context.Background(), rrc21IssuerAddr, nil)
	if err != nil || balance.Cmp(simulation.MinApply) != 0 {
		log.Fatal("can't get balance  in rrc21Issuer SMC: ", err, "got", balance, "wanted", simulation.MinApply)
	}

	//check balance fee
	balanceIssuerFee, err := rrc21Issuer.GetTokenCapacity(rrc21TokenAddr)
	if err != nil || balanceIssuerFee.Cmp(simulation.MinApply) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", simulation.MinApply)
	}
}
