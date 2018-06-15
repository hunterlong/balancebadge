package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func BitcoinBalance(address string) float64 {
	var url string
	url = fmt.Sprintf(BTCapi+"/api/addr/%v/balance", address)
	resp, err := httpGet(url, "GET", []byte(""))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	amount, _ := strconv.ParseFloat(string(body), 10)
	amount = amount * 0.00000001
	return amount
}


