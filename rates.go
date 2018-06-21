package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

var (
	USDrates []*MarketRate
)

func FindCoinRate(coin string) *MarketRate {
	for _, v := range USDrates {
		if coin == v.Symbol {
			return v
		}
	}
	return nil
}

func FetchCoinMarketCap() {
	pages := 15
	var limit int
	currency := "USD"
	var tempRates []*MarketRate
	for i := 0; i <= pages; i++ {
		url := fmt.Sprintf("https://api.coinmarketcap.com/v2/ticker/?start=%v&convert=%v", limit, currency)
		res, err := httpGet(url, "GET", nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Println(err)
			return
		}
		var coinrates CoinMarketCapResponse
		err = json.Unmarshal(body, &coinrates)
		var tickers []*Ticker
		for _, v := range coinrates.Data {
			tickers = append(tickers, v)
		}
		for _, t := range tickers {
			var price float64
			for _, r := range t.Quotes {
				price = r.Price
			}
			rate := &MarketRate{
				Coin:   t.Name,
				Symbol: t.Symbol,
				Price:  price,
			}
			tempRates = append(tempRates, rate)
		}
		limit += 101
		time.Sleep(1 * time.Second)
	}
	USDrates = tempRates
	fmt.Println("Total Coin Market Cap USD Rates: ", len(USDrates))
}
