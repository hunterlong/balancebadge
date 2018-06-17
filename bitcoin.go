package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

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
	return fmt.Sprintf("%f", amount), err
}
