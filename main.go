package main

import (
	"bytes"
	"fmt"
	"github.com/GeertJohan/go.rice"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	BTCapi     string
	BTCTESTapi string
	LTCapi     string
	LTCTESTapi string
	ETHapi     string
	ROPSTENapi string
	svgBox     *rice.Box
	svgData    string
)

func GetEnv() {
	BTCapi = os.Getenv("BTC")
	BTCTESTapi = os.Getenv("BTCTEST")
	LTCapi = os.Getenv("LTC")
	LTCTESTapi = os.Getenv("LTCTEST")
	ETHapi = os.Getenv("ETH")
	ROPSTENapi = os.Getenv("ROPSTEN")
}

func LoadEthBlockchains() error {
	var err error
	eth, err = ethclient.Dial(ETHapi)
	if err != nil {
		panic(err)
	}
	ropsten, err = ethclient.Dial(ROPSTENapi)
	if err != nil {
		panic(err)
	}
	return err
}

func main() {
	GetEnv()
	err := LoadEthBlockchains()
	if err != nil {
		panic(err)
	}

	svgBox = rice.MustFindBox("svg")
	svgData, err = svgBox.String("svg.xml")
	if err != nil {
		fmt.Println(err)
	}

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
	r.Handle("/{coin}/{address}.svg", http.HandlerFunc(NormalBadgeHandler))
	r.Handle("/token/{token}/{address}.svg", http.HandlerFunc(TokenBadgeHandler))
	return r
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://docs.balancebadge.io", http.StatusSeeOther)
}

func CryptoBalance(coin, address string) string {
	var balance string
	var err error
	switch coin {
	case "btc":
		balance, err = BitcoinBalance(BTCapi, address)
		if err != nil {
			fmt.Println(err)
			return "0"
		}
	case "btctest":
		balance, err = BitcoinBalance(BTCTESTapi, address)
		if err != nil {
			fmt.Println(err)
			return "0"
		}
	case "ltc":
		balance, err = BitcoinBalance(LTCapi, address)
		if err != nil {
			return "0"
		}
	case "ltctest":
		balance, err = BitcoinBalance(LTCTESTapi, address)
		if err != nil {
			return "0"
		}
	case "eth":
		balance = EthBalance(eth, address)
	case "ropsten":
		balance = EthBalance(ropsten, address)
	}
	return balance
}

type Badge struct {
	Coin          string
	Token         string
	Address       string
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
}

func (b *Badge) Clean() *Badge {
	b.LeftSize = len(b.Label) * 9
	b.LeftTextSize = (b.LeftSize * 9) - 65
	b.LeftTextX = (b.LeftTextSize / 2) + 60

	//leftSum := b.LeftTextSize + b.LeftSize + b.LeftTextX

	fmt.Println("LEFT  ", b.LeftSize, b.LeftTextSize, b.LeftTextX)

	b.RightSize = (len(b.Balance) + len(b.Coin)) * 10
	b.RightTextSize = b.RightSize * 8
	b.RightTextX = b.LeftTextSize + 510

	fmt.Println("RIGHT ", b.RightSize, b.RightTextSize, b.RightTextX)

	if b.RightSize < 75 {
		b.RightSize = 75
	}
	//if b.RightTextSize < 640 {
	//	b.RightTextSize = 640
	//}
	//
	//if b.RightTextX < 980 {
	//	b.RightTextX = 980
	//}
	//if b.RightTextX > 1090 {
	//	b.RightTextX = 1090
	//}
	b.Width = b.LeftSize + b.RightSize
	return b
}

func (b *Badge) Normal() *Badge {
	if b.Token == "" {
		b.Balance = CryptoBalance(b.Coin, b.Address)
	}
	rightColor := "97CA00"

	if b.RightColor != "" {
		rightColor = b.RightColor
	}

	label := b.Address[0:7]
	if b.Label != "" {
		label = b.Label
	}

	badge := &Badge{
		Coin:       strings.ToUpper(b.Coin),
		Address:    b.Address[0:7],
		Balance:    b.Balance,
		Label:      label,
		Type:       b.Type,
		Height:     20,
		LeftColor:  "555555",
		RightColor: rightColor,
	}
	return badge.Clean()
}

func TokenBadgeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	token, _ := vars["token"]
	address, _ := vars["address"]
	badgeType, _ := vars["type"]
	color := r.FormValue("color")
	label := r.FormValue("label")
	temp := template.New("svg")
	temp.Parse(string(svgData))

	balance, symbol, err := TokenBalance(token, address)
	if err != nil {
		panic(err)
	}

	badge := &Badge{
		Coin:       symbol,
		Token:      token,
		Balance:    balance,
		Address:    address,
		Label:      label,
		Type:       badgeType,
		RightColor: color,
	}

	badgeSvg := badge.Normal()

	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	temp.Execute(w, badgeSvg)
}

func NormalBadgeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin, _ := vars["coin"]
	address, _ := vars["address"]
	badgeType, _ := vars["type"]
	color := r.FormValue("color")
	label := r.FormValue("label")

	temp := template.New("svg")
	temp.Parse(string(svgData))

	badge := &Badge{
		Coin:       coin,
		Address:    address,
		Label:      label,
		Type:       badgeType,
		RightColor: color,
	}

	badgeSvg := badge.Normal()

	w.Header().Set("Content-Type", "image/svg+xml")
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	temp.Execute(w, badgeSvg)
}

func httpGet(url string, method string, data []byte) (*http.Response, error) {
	req, _ := http.NewRequest(method, url, bytes.NewBuffer(data))
	httpClient := &http.Client{}
	resp, err := httpClient.Do(req)
	return resp, err
}
