package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"fmt"
	"bytes"
	"os"
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
	r.Handle("/{coin}/{address}/{type}", http.HandlerFunc(BadgeHandler))
	r.Handle("/{coin}/{address}/{type}/usd", http.HandlerFunc(BadgeHandler))
	return r
}


func IndexHandler(w http.ResponseWriter, r *http.Request) {
	badge := []byte(`<svg xmlns="http://www.w3.org/2000/svg" width="104" height="20"><linearGradient id="b" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><mask id="a"><rect width="104" height="20" rx="3" fill="#fff"/></mask><g mask="url(#a)"><path fill="#555" d="M0 0h54v20H0z"/><path fill="#4c1" d="M54 0h50v20H54z"/><path fill="url(#b)" d="M0 0h104v20H0z"/></g><g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11"><text x="28" y="15" fill="#010101" fill-opacity=".3">h7BowIw</text><text x="28" y="14">h7BowIw</text><text x="78" y="15" fill="#010101" fill-opacity=".3">12.533</text><text x="78" y="14">12.533</text></g></svg>`)
	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Write(badge)
}

func BadgeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin, _ := vars["coin"]
	address, _ := vars["address"]
	badgeType, _ := vars["type"]

	var balance float64

	switch coin {
		case "btc": balance = BitcoinBalance(address)
	}
	if balance == 0 {
		balance = 0.0
	}

	fmt.Println("balance: ", balance)

	fmt.Println(badgeType, coin)

	addressFormat := address[0:6]

	badgeData := fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="99" height="20"><linearGradient id="b" x2="0" y2="100%%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient><mask id="a"><rect width="99" height="20" rx="3" fill="#fff"/></mask><g mask="url(#a)"><path fill="#555" d="M0 0h54v20H0z"/><path fill="#e05d44" d="M54 0h45v20H54z"/><path fill="url(#b)" d="M0 0h99v20H0z"/></g><g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="11"><text x="28" y="15" fill="#010101" fill-opacity=".3">%v</text><text x="28" y="14">%v</text><text x="75.5" y="15" fill="#010101" fill-opacity=".3">%0.3f %v</text><text x="75.5" y="14">%0.3f %v</text></g></svg>`, addressFormat, addressFormat, balance, coin, balance, coin)

	badge := []byte(badgeData)

	WriteBadge(badge, w)
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