package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func BitcoinBalance(address string) (float64, error) {
	var url string
	url = fmt.Sprintf(BTCapi+"/api/addr/%v/balance", address)
	resp, err := httpGet(url, "GET", []byte(""))
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	amount, err := strconv.ParseFloat(string(body), 10)
	if err != nil {
		return 0, err
	}
	amount = amount * 0.00000001
	return amount, err
}
