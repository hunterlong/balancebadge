package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func CryptoBalance(coin, address string) string {
	var balance string
	var err error
	switch coin {
	case "BTC":
		balance, err = BitcoinBalance(BTCapi, address)
		if err != nil {
			fmt.Println(err)
			return "0"
		}
	case "BCH":
		balance, err = BitcoinBalance(BCHapi, address)
		if err != nil {
			fmt.Println(err)
			return "0"
		}
	case "BCHTEST":
		balance, err = BitcoinBalance(BCHTESTapi, address)
		if err != nil {
			fmt.Println(err)
			return "0"
		}
	case "BTCTEST":
		balance, err = BitcoinBalance(BTCTESTapi, address)
		if err != nil {
			fmt.Println(err)
			return "0"
		}
	case "LTC":
		balance, err = BitcoinBalance(LTCapi, address)
		if err != nil {
			return "0"
		}
	case "LTCTEST":
		balance, err = BitcoinBalance(LTCTESTapi, address)
		if err != nil {
			return "0"
		}
	case "DASH":
		balance, err = BitcoinBalance(DASHapi, address)
		if err != nil {
			return "0"
		}
	case "ZCASH":
		balance, err = BitcoinBalance(ZCASHapi, address)
		if err != nil {
			return "0"
		}
	case "ETH":
		balance = EthBalance(eth, address)
	case "ROPSTEN":
		balance = EthBalance(ropsten, address)
	case "RINKEBY":
		balance = EthBalance(rinkeby, address)
	default:
		balance = "0"
	}
	return balance
}

func BitcoinBalance(endpoint, address string) (string, error) {
	var url string
	url = fmt.Sprintf(endpoint+"/addr/%v/balance", address)
	resp, err := httpGet(url, "GET", []byte(""))
	if err != nil {
		return "0", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "0", err
	}
	amount, err := strconv.ParseFloat(string(body), 10)
	if err != nil {
		return "0", err
	}
	amount = amount * 0.00000001
	return fmt.Sprintf("%0.3f", amount), err
}
