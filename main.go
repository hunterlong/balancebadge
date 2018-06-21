package main

import (
	"time"
)

var (
	BTCapi           string
	BTCTESTapi       string
	LTCapi           string
	LTCTESTapi       string
	ETHapi           string
	ROPSTENapi       string
	status24HourHits int64
)

func main() {
	GetEnv()
	err := LoadEthBlockchains()
	if err != nil {
		panic(err)
	}
	go CoinMarketCapTicker()
	go HourLoop()
	StartHTTPServer()
}

func HourLoop() {
	defer HourLoop()
	status24HourHits = 0
	time.Sleep(60 * time.Minute)
}
