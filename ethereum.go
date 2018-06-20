package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"io/ioutil"
	"math/big"
	"strconv"
)

var (
	eth     *ethclient.Client
	ropsten *ethclient.Client
)

func EthBalance(api *ethclient.Client, address string) string {
	balance, _ := api.BalanceAt(context.TODO(), common.HexToAddress(address), nil)
	amount := BigIntDecimal(balance, 18)
	amountFloat, _ := strconv.ParseFloat(amount, 10)
	return fmt.Sprintf("%0.3f", amountFloat)
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

func TokenBalance(token, address string) (string, string, error) {
	var url string
	url = fmt.Sprintf("https://api.tokenbalance.com/token/%v/%v", token, address)
	resp, err := httpGet(url, "GET", []byte(""))
	if err != nil {
		return "0", "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "0", "", err
	}

	var tokenBalance TokenBalanceResponse
	json.Unmarshal(body, &tokenBalance)

	amount, err := strconv.ParseFloat(tokenBalance.Balance, 10)
	if err != nil {
		return "0", "", err
	}

	return fmt.Sprintf("%0.3f", amount), tokenBalance.Symbol, err
}

type TokenBalanceResponse struct {
	Name       string `json:"name"`
	Wallet     string `json:"wallet"`
	Symbol     string `json:"symbol"`
	Balance    string `json:"balance"`
	EthBalance string `json:"eth_balance"`
	Decimals   int    `json:"decimals"`
	Block      int    `json:"block"`
}
