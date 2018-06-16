package main

import (
	"net/http"
	"github.com/gorilla/mux"
	"time"
	"fmt"
	"bytes"
	"os"
	"strings"
	"html/template"
	"io/ioutil"
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

	file, _ := ioutil.ReadFile("svg.xml")

	fmt.Println(file)

	temp := template.New("svg")
	temp.Parse(string(file))

	bb := &Badge{
		Coin: "btc",
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
		if err != nil { return 0 }
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
	Width string
	Height string
	LeftColor string
	LeftSize string
	RightColor string
	RightSize string
}


func (b *Badge) Normal() string {
	leftWidth := "60"
	rightWidth := "120"
	fullWidth := leftWidth + rightWidth
	balance := CryptoBalance(b.Coin, b.Address)
	addressFormat := b.Address[0:7]
	coin := strings.ToUpper(b.Coin)

	leftColor := "#555555"
	rightColor := "#97CA00"

	clip := fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="%v" height="20"><linearGradient id="b" x2="0" y2="100%%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/><stop offset="1" stop-opacity=".1"/></linearGradient>`, fullWidth)
	clip += fmt.Sprintf(`<clipPath id="a"><rect width="134" height="20" rx="3" fill="#fff"/></clipPath>`)
	clip += fmt.Sprintf(`<g clip-path="url(#a)"><path fill="%v" d="M0 0h%vv20H0z"/><path fill="%v" d="M%v 0h%vv20H%vz"/><path fill="url(#b)" d="M0 0h%vv20H0z"/></g>`, leftColor, leftWidth, rightColor, leftWidth, rightWidth, leftWidth, fullWidth)
	clip += fmt.Sprintf(`<g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="110">`)
	clip += fmt.Sprintf(`<text x="305" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="490">%v</text>`, addressFormat)
	clip += fmt.Sprintf(`<text x="305" y="140" transform="scale(.1)" textLength="490">%v</text>`, addressFormat)
	clip += fmt.Sprintf(`<text x="955" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="650">%0.3f %v</text>`, balance, coin)
	clip += fmt.Sprintf(`<text x="955" y="140" transform="scale(.1)" textLength="650">%0.3f %v</text>`, balance, coin)
	clip += fmt.Sprintf(`</g></svg>`)
	return clip
}


func NormalBadgeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin, _ := vars["coin"]
	address, _ := vars["address"]
	badgeType, _ := vars["type"]

	badge := &Badge{
		Coin: coin,
		Address: address,
		Type: badgeType,
	}

	fmt.Println(badge)

	badgeSvg := badge.Normal()

	badgeD := []byte(badgeSvg)

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