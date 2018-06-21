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
