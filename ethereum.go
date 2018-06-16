package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strconv"
)

var (
	eth *ethclient.Client
)

func EthBalance(address string) string {
	balance, _ := eth.BalanceAt(context.TODO(), common.HexToAddress(address), nil)
	return BigIntDecimal(balance, 18)
}

func BigIntDecimal(balance *big.Int, decimals int64) string {
	if balance.Sign() == 0 {
		return "0"
	}
	bal := big.NewFloat(0)
	bal.SetInt(balance)
	pow := BigPow(10, decimals)
	p := big.NewFloat(0)
	p.SetInt(pow)
	bal.Quo(bal, p)
	deci := strconv.Itoa(int(decimals))
	dec := "%." + deci + "f"
	newNum := Clean(fmt.Sprintf(dec, bal))
	return newNum
}

func Clean(newNum string) string {
	stringBytes := bytes.TrimRight([]byte(newNum), "0")
	newNum = string(stringBytes)
	if stringBytes[len(stringBytes)-1] == 46 {
		newNum += "0"
	}
	if stringBytes[0] == 46 {
		newNum = "0" + newNum
	}
	return newNum
}

func BigPow(a, b int64) *big.Int {
	r := big.NewInt(a)
	return r.Exp(r, big.NewInt(b), nil)
}
