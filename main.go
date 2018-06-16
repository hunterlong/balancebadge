package main

import (
	"bytes"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	BTCapi string
)

func GetEnv() {
	BTCapi = os.Getenv("BTC")
}

func init() {
	var err error
	GetEnv()
	eth, err = ethclient.Dial("https://eth.coinapp.io")
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("BALANCE BADGE running on http://localhost:9090")
	r := Router()
	srv := &http.Server{
		Addr:         "0.0.0.0:9090",
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		IdleTimeout:  time.Second * 30,
		Handler:      r,
	}
	srv.ListenAndServe()
}

func Router() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/", http.HandlerFunc(IndexHandler))
	r.Handle("/{coin}/{address}", http.HandlerFunc(NormalBadgeHandler))
	return r
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	file, _ := ioutil.ReadFile("svg.xml")

	fmt.Println(file)

	temp := template.New("svg")
	temp.Parse(string(file))

	bb := &Badge{
		Coin:    "btc",
		Address: "0x0x0x0x00x0x0x0x0x0",
	}

	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	temp.Execute(w, bb)
}

func CryptoBalance(coin, address string) float64 {
	var balance float64
	var err error
	switch coin {
	case "btc":
		balance, err = BitcoinBalance(address)
		if err != nil {
			return 0
		}
	}
	if balance == 0 {
		balance = 0.0
	}
	return balance
}

type Badge struct {
	Coin       string
	Address    string
	Balance    string
	Type       string
	Width      int
	Height     int
	LeftColor  string
	LeftSize   int
	RightColor string
	RightSize  int
}

func (b *Badge) Normal() *Badge {
	var balance float64
	if b.Coin == "eth" {
		ethBal := EthBalance(b.Address)
		balance, _ = strconv.ParseFloat(ethBal, 10)
	} else {
		balance = CryptoBalance(b.Coin, b.Address)
	}

	badge := &Badge{
		Coin:       strings.ToUpper(b.Coin),
		Address:    b.Address[0:7],
		Balance:    fmt.Sprintf("%0.3f", balance),
		Type:       b.Type,
		Height:     20,
		LeftColor:  "#555555",
		LeftSize:   60,
		RightColor: "#97CA00",
		RightSize:  75,
	}
	badge.Width = badge.LeftSize + badge.RightSize
	return badge
}

func NormalBadgeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin, _ := vars["coin"]
	address, _ := vars["address"]
	badgeType, _ := vars["type"]

	file, _ := ioutil.ReadFile("svg.xml")
	temp := template.New("svg")
	temp.Parse(string(file))

	badge := &Badge{
		Coin:    coin,
		Address: address,
		Type:    badgeType,
	}

	badgeSvg := badge.Normal()

	fmt.Println(badgeSvg)

	//w.Header().Set("Content-Type", "image/svg+xml")
	//w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")

	temp.Execute(w, badgeSvg)
}

func WriteBadge(badge []byte, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Write(badge)
}

func httpGet(url string, method string, data []byte) (*http.Response, error) {
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(data))
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	return resp, err
}
