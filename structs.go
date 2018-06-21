package main

import "time"

type Server struct {
	Online     bool  `json:"online"`
	Hour24Hits int64 `json:"24_hour_requests"`
}

type Badge struct {
	Coin          string
	Token         string
	Address       string
	FullAddress   string
	Balance       string
	Label         string
	Type          string
	Width         int
	Height        int
	LeftColor     string
	LeftSize      int
	LeftTextSize  int
	LeftTextX     int
	RightColor    string
	RightSize     int
	RightTextSize int
	RightTextX    int
	error         error
	start         time.Time
	toCurrency    bool
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

type MarketRate struct {
	Coin   string
	Symbol string
	Price  float64
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
