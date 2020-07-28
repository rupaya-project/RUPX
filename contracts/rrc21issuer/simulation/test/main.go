package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/rupaya-project/rupx/accounts/abi/bind"
	"github.com/rupaya-project/rupx/common"
	"github.com/rupaya-project/rupx/common/hexutil"
	rrc21issuer "github.com/rupaya-project/rupx/contracts/rrc21issuer"
	"github.com/rupaya-project/rupx/contracts/rrc21issuer/simulation"
	"github.com/rupaya-project/rupx/ethclient"
)

var (
	rrc21TokenAddr = common.HexToAddress("0x80430A33EaB86890a346bCf64F86CFeAC73287f3")
)

func airDropTokenToAccountNoRupaya() {
	client, err := ethclient.Dial(simulation.RpcEndpoint)
	if err != nil {
		fmt.Println(err, client)
	}
	nonce, _ := client.NonceAt(context.Background(), simulation.MainAddr, nil)
	mainAccount := bind.NewKeyedTransactor(simulation.MainKey)
	mainAccount.Nonce = big.NewInt(int64(nonce))
	mainAccount.Value = big.NewInt(0)      // in wei
	mainAccount.GasLimit = uint64(4000000) // in units
	mainAccount.GasPrice = big.NewInt(0).Mul(common.RRC21GasPrice, big.NewInt(2))
	rrc21Instance, _ := rrc21issuer.NewRRC21(mainAccount, rrc21TokenAddr, client)
	rrc21IssuerInstance, _ := rrc21issuer.NewRRC21Issuer(mainAccount, common.RRC21IssuerSMC, client)
	// air drop token
	remainFee, _ := rrc21IssuerInstance.GetTokenCapacity(rrc21TokenAddr)
	tx, err := rrc21Instance.Transfer(simulation.AirdropAddr, simulation.AirDropAmount)
	if err != nil {
		log.Fatal("can't air drop to ", err)
	}
	// check balance after transferAmount
	fmt.Println("wait 10s to airdrop success ", tx.Hash().Hex())
	time.Sleep(10 * time.Second)

	_, receiptRpc, err := client.GetTransactionReceiptResult(context.Background(), tx.Hash())
	receipt := map[string]interface{}{}
	err = json.Unmarshal(receiptRpc, &receipt)
	if err != nil {
		log.Fatal("can't transaction's receipt ", err, "hash", tx.Hash().Hex())
	}
	fee := big.NewInt(0).SetUint64(hexutil.MustDecodeUint64(receipt["gasUsed"].(string)))
	if hexutil.MustDecodeUint64(receipt["blockNumber"].(string)) > common.TIPRRC21Fee.Uint64() {
		fee = fee.Mul(fee, common.RRC21GasPrice)
	}
	fmt.Println("fee", fee.Uint64(), "number", hexutil.MustDecodeUint64(receipt["blockNumber"].(string)))
	remainFee = big.NewInt(0).Sub(remainFee, fee)
	//check balance fee
	balanceIssuerFee, err := rrc21IssuerInstance.GetTokenCapacity(rrc21TokenAddr)
	if err != nil || balanceIssuerFee.Cmp(remainFee) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
	if err != nil {
		log.Fatal("can't execute transferAmount in tr21:", err)
	}
}
func testTransferRRC21TokenWithAccountNoRupaya() {
	client, err := ethclient.Dial(simulation.RpcEndpoint)
	if err != nil {
		fmt.Println(err, client)
	}

	// access to address which received token rrc20 but dont have rupaya
	nonce, _ := client.NonceAt(context.Background(), simulation.AirdropAddr, nil)
	airDropAccount := bind.NewKeyedTransactor(simulation.AirdropKey)
	airDropAccount.Nonce = big.NewInt(int64(nonce))
	airDropAccount.Value = big.NewInt(0)      // in wei
	airDropAccount.GasLimit = uint64(4000000) // in units
	airDropAccount.GasPrice = big.NewInt(0).Mul(common.RRC21GasPrice, big.NewInt(2))
	rrc21Instance, _ := rrc21issuer.NewRRC21(airDropAccount, rrc21TokenAddr, client)
	rrc21IssuerInstance, _ := rrc21issuer.NewRRC21Issuer(airDropAccount, common.RRC21IssuerSMC, client)

	remainFee, _ := rrc21IssuerInstance.GetTokenCapacity(rrc21TokenAddr)
	airDropBalanceBefore, err := rrc21Instance.BalanceOf(simulation.AirdropAddr)
	receiverBalanceBefore, err := rrc21Instance.BalanceOf(simulation.ReceiverAddr)
	// execute transferAmount trc to other address
	tx, err := rrc21Instance.Transfer(simulation.ReceiverAddr, simulation.TransferAmount)
	if err != nil {
		log.Fatal("can't execute transferAmount in tr21:", err)
	}

	// check balance after transferAmount
	fmt.Println("wait 10s to transferAmount success ")
	time.Sleep(10 * time.Second)

	balance, err := rrc21Instance.BalanceOf(simulation.ReceiverAddr)
	wantedBalance := big.NewInt(0).Add(receiverBalanceBefore, simulation.TransferAmount)
	if err != nil || balance.Cmp(wantedBalance) != 0 {
		log.Fatal("check balance after fail receiverAmount in tr21: ", err, "get", balance, "wanted", wantedBalance)
	}

	remainAirDrop := big.NewInt(0).Sub(airDropBalanceBefore, simulation.TransferAmount)
	remainAirDrop = remainAirDrop.Sub(remainAirDrop, simulation.Fee)
	// check balance rrc21 again
	balance, err = rrc21Instance.BalanceOf(simulation.AirdropAddr)
	if err != nil || balance.Cmp(remainAirDrop) != 0 {
		log.Fatal("check balance after fail transferAmount in tr21: ", err, "get", balance, "wanted", remainAirDrop)
	}
	_, receiptRpc, err := client.GetTransactionReceiptResult(context.Background(), tx.Hash())
	receipt := map[string]interface{}{}
	err = json.Unmarshal(receiptRpc, &receipt)
	if err != nil {
		log.Fatal("can't transaction's receipt ", err, "hash", tx.Hash().Hex())
	}
	fee := big.NewInt(0).SetUint64(hexutil.MustDecodeUint64(receipt["gasUsed"].(string)))
	if hexutil.MustDecodeUint64(receipt["blockNumber"].(string)) > common.TIPRRC21Fee.Uint64() {
		fee = fee.Mul(fee, common.RRC21GasPrice)
	}
	fmt.Println("fee", fee.Uint64(), "number", hexutil.MustDecodeUint64(receipt["blockNumber"].(string)))
	remainFee = big.NewInt(0).Sub(remainFee, fee)
	//check balance fee
	balanceIssuerFee, err := rrc21IssuerInstance.GetTokenCapacity(rrc21TokenAddr)
	if err != nil || balanceIssuerFee.Cmp(remainFee) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
	//check rrc21 SMC balance
	balance, err = client.BalanceAt(context.Background(), common.RRC21IssuerSMC, nil)
	if err != nil || balance.Cmp(remainFee) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
}
func testTransferTrc21Fail() {
	client, err := ethclient.Dial(simulation.RpcEndpoint)
	if err != nil {
		fmt.Println(err, client)
	}
	nonce, _ := client.NonceAt(context.Background(), simulation.AirdropAddr, nil)
	airDropAccount := bind.NewKeyedTransactor(simulation.AirdropKey)
	airDropAccount.Nonce = big.NewInt(int64(nonce))
	airDropAccount.Value = big.NewInt(0)      // in wei
	airDropAccount.GasLimit = uint64(4000000) // in units
	airDropAccount.GasPrice = big.NewInt(0).Mul(common.RRC21GasPrice, big.NewInt(2))
	rrc21Instance, _ := rrc21issuer.NewRRC21(airDropAccount, rrc21TokenAddr, client)
	rrc21IssuerInstance, _ := rrc21issuer.NewRRC21Issuer(airDropAccount, common.RRC21IssuerSMC, client)
	balanceIssuerFee, err := rrc21IssuerInstance.GetTokenCapacity(rrc21TokenAddr)

	minFee, err := rrc21Instance.MinFee()
	if err != nil {
		log.Fatal("can't get minFee of rrc21 smart contract:", err)
	}
	ownerBalance, err := rrc21Instance.BalanceOf(simulation.MainAddr)
	remainFee, err := rrc21IssuerInstance.GetTokenCapacity(rrc21TokenAddr)
	airDropBalanceBefore, err := rrc21Instance.BalanceOf(simulation.AirdropAddr)

	tx, err := rrc21Instance.Transfer(common.Address{}, big.NewInt(1))
	if err != nil {
		log.Fatal("can't execute test transfer to zero address in tr21:", err)
	}
	fmt.Println("wait 10s to transfer to zero address")
	time.Sleep(10 * time.Second)

	fmt.Println("airDropBalanceBefore", airDropBalanceBefore)
	// check balance rrc21 again
	airDropBalanceBefore = big.NewInt(0).Sub(airDropBalanceBefore, minFee)
	balance, err := rrc21Instance.BalanceOf(simulation.AirdropAddr)
	if err != nil || balance.Cmp(airDropBalanceBefore) != 0 {
		log.Fatal("check balance after fail transferAmount in tr21: ", err, "get", balance, "wanted", airDropBalanceBefore)
	}

	ownerBalance = big.NewInt(0).Add(ownerBalance, minFee)
	//check balance fee
	balance, err = rrc21Instance.BalanceOf(simulation.MainAddr)
	if err != nil || balance.Cmp(ownerBalance) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
	_, receiptRpc, err := client.GetTransactionReceiptResult(context.Background(), tx.Hash())
	receipt := map[string]interface{}{}
	err = json.Unmarshal(receiptRpc, &receipt)
	if err != nil {
		log.Fatal("can't transaction's receipt ", err, "hash", tx.Hash().Hex())
	}
	fee := big.NewInt(0).SetUint64(hexutil.MustDecodeUint64(receipt["gasUsed"].(string)))
	if hexutil.MustDecodeUint64(receipt["blockNumber"].(string)) > common.TIPRRC21Fee.Uint64() {
		fee = fee.Mul(fee, common.RRC21GasPrice)
	}
	fmt.Println("fee", fee.Uint64(), "number", hexutil.MustDecodeUint64(receipt["blockNumber"].(string)))
	remainFee = big.NewInt(0).Sub(remainFee, fee)
	//check balance fee
	balanceIssuerFee, err = rrc21IssuerInstance.GetTokenCapacity(rrc21TokenAddr)
	if err != nil || balanceIssuerFee.Cmp(remainFee) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}
	//check rrc21 SMC balance
	balance, err = client.BalanceAt(context.Background(), common.RRC21IssuerSMC, nil)
	if err != nil || balance.Cmp(remainFee) != 0 {
		log.Fatal("can't get balance token fee in  smart contract: ", err, "got", balanceIssuerFee, "wanted", remainFee)
	}

}
func main() {
	fmt.Println("========================")
	fmt.Println("airdropAddr", simulation.AirdropAddr.Hex())
	fmt.Println("receiverAddr", simulation.ReceiverAddr.Hex())
	fmt.Println("========================")

	start := time.Now()
	for i := 0; i < 10000000; i++ {
		airDropTokenToAccountNoRupaya()
		fmt.Println("Finish airdrop token to a account")
		testTransferRRC21TokenWithAccountNoRupaya()
		fmt.Println("Finish transfer rrc21 token with a account no rupaya")
		testTransferTrc21Fail()
		fmt.Println("Finish testing ! Success transferAmount token rrc20 with a account no rupaya")
	}
	fmt.Println(common.PrettyDuration(time.Since(start)))
}
