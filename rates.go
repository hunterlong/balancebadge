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

type MarketRate struct {
	Coin   string
	Symbol string
	Price  float64
}

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
			USDrates = append(USDrates, rate)
		}
		limit += 101
		time.Sleep(1 * time.Second)
	}
	fmt.Println("Total Coin Market Cap USD Rates: ", len(USDrates))
}

type CoinMarketCapResponse struct {
	Data     map[string]*Ticker `json:"data,omitempty"`
	Metadata struct {
		Timestamp           int64
		NumCryptoCurrencies int    `json:"num_cryptocurrencies,omitempty"`
		Error               string `json:",omitempty"`
	}
}

type Ticker struct {
	ID                int                     `json:"id"`
	Name              string                  `json:"name"`
	Symbol            string                  `json:"symbol"`
	Slug              string                  `json:"website_slug"`
	Rank              int                     `json:"rank"`
	CirculatingSupply float64                 `json:"circulating_supply"`
	TotalSupply       float64                 `json:"total_supply"`
	MaxSupply         float64                 `json:"max_supply"`
	Quotes            map[string]*TickerQuote `json:"quotes"`
	LastUpdated       int                     `json:"last_updated"`
}

type TickerQuote struct {
	Price            float64 `json:"price"`
	Volume24H        float64 `json:"volume_24h"`
	MarketCap        float64 `json:"market_cap"`
	PercentChange1H  float64 `json:"percent_change_1h"`
	PercentChange24H float64 `json:"percent_change_24h"`
	PercentChange7D  float64 `json:"percent_change_7d"`
}
