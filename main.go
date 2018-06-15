package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"fmt"
	"bytes"
	"os"
	"strings"
)

var (
	BTCapi string
)


func GetEnv() {
	BTCapi = os.Getenv("BTC")
}


func main() {
	GetEnv()
	fmt.Println("CRYBADGE running on http://localhost:9090")
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
	badge := []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="1" height="20"><linearGradient id="b" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><mask id="a"><rect width="120" height="20" rx="3" fill="#fff"/></mask><g mask="url(#a)"><path fill="#555" d="M0 0h80v20H0z"/><path fill="#4c1" d="M54 0h80v20H54z"/><path fill="url(#b)" d="M0 0h130v20H0z"/></g><g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11"><text x="28" y="15" fill="#010101" fill-opacity=".3">h7BowIw</text><text x="28" y="14">h7BowIw</text><text x="78" y="15" fill="#010101" fill-opacity=".3">12.533</text><text x="78" y="14">12.533</text></g></svg>`)
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Write(badge)
}


func CryptoBalance(coin, address string) float64 {
	var balance float64
	switch coin {
	case "btc": balance = BitcoinBalance(address)
	}
	if balance == 0 {
		balance = 0.0
	}
	return balance
}

type Badge struct {
	Coin string
	Address string
	Type string
}

func NormalBadgeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin, _ := vars["coin"]
	address, _ := vars["address"]
	badgeType, _ := vars["type"]

	badge := &Badge{
		coin,
		address,
		badgeType,
	}

	fmt.Println(badge)

	balance := CryptoBalance(coin, address)

	fmt.Println(badgeType, coin)

	addressFormat := address[0:7]

	coin = strings.ToUpper(coin)

	badgeData := fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="134" height="20"><linearGradient id="b" x2="0" y2="100%%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><clipPath id="a"><rect width="134" height="20" rx="3" fill="#fff"/></clipPath><g clip-path="url(#a)"><path fill="#555" d="M0 0h59v20H0z"/><path fill="#97CA00" d="M59 0h75v20H59z"/><path fill="url(#b)" d="M0 0h134v20H0z"/></g><g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="110"><text x="305" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="490">%v</text><text x="305" y="140" transform="scale(.1)" textLength="490">%v</text><text x="955" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="650">%0.3f %v</text><text x="955" y="140" transform="scale(.1)" textLength="650">%0.3f %v</text></g> </svg>`, addressFormat, addressFormat, balance, coin, balance, coin)

	badgeD := []byte(badgeData)

	WriteBadge(badgeD, w)
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