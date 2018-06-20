package main

import (
	"bytes"
	"fmt"
	"github.com/GeertJohan/go.rice"
	cmc "github.com/coincircle/go-coinmarketcap"
	"github.com/coincircle/go-coinmarketcap/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
	"strconv"
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
	tickers    []*types.Ticker
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

	CoinMarketCapTicker()

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

type TickerObject struct {
	Tickers []Ticker `json:"data"`
}

type Ticker struct {
	ID                int         `json:"id"`
	Name              string      `json:"name"`
	Symbol            string      `json:"symbol"`
	WebsiteSlug       string      `json:"website_slug"`
	Rank              int         `json:"rank"`
	CirculatingSupply float64     `json:"circulating_supply"`
	TotalSupply       float64     `json:"total_supply"`
	MaxSupply         interface{} `json:"max_supply"`
	Quotes            struct {
		USD struct {
			Price            float64 `json:"price"`
			Volume24H        float64 `json:"volume_24h"`
			MarketCap        float64 `json:"market_cap"`
			PercentChange1H  float64 `json:"percent_change_1h"`
			PercentChange24H float64 `json:"percent_change_24h"`
			PercentChange7D  float64 `json:"percent_change_7d"`
		} `json:"USD"`
	} `json:"quotes"`
	LastUpdated int `json:"last_updated"`
}

func CoinMarketCapTicker() {
	var err error
	tickers, err = cmc.Tickers(&cmc.TickersOptions{
		Start:   0,
		Limit:   10000,
		Convert: "USD",
	})
	if err != nil {
		panic(err)
	}

	btcs := CoinToUSD("btc")

	fmt.Println(btcs)
}

func CoinToUSD(coin string) float64 {
	upper := strings.ToUpper(coin)
	for _, ticker := range tickers {
		if ticker.Symbol == upper {
			return ticker.Quotes["USD"].Price
		}
	}
	return 0
}

func Router() *mux.Router {
	r := mux.NewRouter()
	r.Handle("/", http.HandlerFunc(IndexHandler))
	r.Handle("/{coin}/{address}.svg", http.HandlerFunc(NormalBadgeHandler))
	r.Handle("/token/{token}/{address}.svg", http.HandlerFunc(TokenBadgeHandler))
	r.Handle("/{coin}/{address}/usd.svg", http.HandlerFunc(USDBadgeHandler))

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

func ColorToHex(color string) string {
	var hex string
	switch color {
	case "green":
		hex = "32c12c"
	case "teal":
		hex = "009888"
	case "indigo":
		hex = "3e49bb"
	case "blue":
		hex = "526eff"
	case "purple":
		hex = "7f4fc9"
	case "lightgreen":
		hex = "87c735"
	case "lime":
		hex = "cde000"
	case "lightblue":
		hex = "00a5f9"
	case "cyan":
		hex = "00bcd9"
	case "darkpurple":
		hex = "682cbf"
	case "yellow":
		hex = "ffef00"
	case "orange":
		hex = "ff9a00"
	case "lightred":
		hex = "ff9a00"
	case "brown":
		hex = "7c5547"
	case "bluegrey":
		hex = "5f7d8e"
	case "amber":
		hex = "ffcd00"
	case "darkorange":
		hex = "ff5500"
	case "red":
		hex = "d40c00"
	case "darkbrown":
		hex = "50342c"
	case "grey":
		hex = "9e9e9e"
	case "white":
		hex = "ffffff"
	case "black":
		hex = "000000"
	default:
		hex = color
	}

	return hex
}

func (b *Badge) Clean() *Badge {
	balanceFloat, _ := strconv.ParseFloat(b.Balance, 10)
	if balanceFloat > 1000 {
		b.Balance = RenderFloat("#,###.", b.Balance)
	} else if balanceFloat > 100 {
		b.Balance = RenderFloat("#,###.#", b.Balance)
	} else if balanceFloat > 10 {
		b.Balance = RenderFloat("#,###.##", b.Balance)
	} else {
		b.Balance = RenderFloat("#,###.###", b.Balance)
	}
	if b.Coin == "USD" {
		b.Balance = "$" + b.Balance
	}

	b.LeftSize = (len(b.Label) + 1) * 8
	b.LeftTextSize = b.LeftSize * 8
	b.LeftTextX = (b.LeftTextSize / 2) + 60

	b.RightSize = (len(b.Balance) + len(b.Coin) + 1) * 8
	b.RightTextSize = (b.RightSize * 8) + 40
	b.RightTextX = (b.RightTextSize / 2) + (b.LeftSize * 2) + b.LeftTextSize + 60

	b.Width = b.LeftSize + b.RightSize

	return b
}

func (b *Badge) Normal() *Badge {
	if b.Token == "" && b.Coin != "USD" {
		b.Balance = CryptoBalance(b.Coin, b.Address)
	}
	rightColor := "97CA00"

	if b.RightColor != "" {
		rightColor = ColorToHex(b.RightColor)
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
	fmt.Println(badge.Balance)
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

func USDBadgeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	coin, _ := vars["coin"]
	address, _ := vars["address"]
	badgeType, _ := vars["type"]
	color := r.FormValue("color")
	label := r.FormValue("label")

	temp := template.New("svg")
	temp.Parse(string(svgData))

	usd := CoinToUSD(coin)

	var balance string
	if coin == "token" {
		balance, _, _ = TokenBalance(coin, address)
	} else {
		balance = CryptoBalance(coin, address)
	}

	fmt.Println(usd, balance)

	balFloat, _ := strconv.ParseFloat(balance, 10)
	usdBalance := usd * balFloat
	usdString := fmt.Sprintf("%0.3f", usdBalance)

	fmt.Printf("%v\n", usdString)

	badge := &Badge{
		Coin:       "USD",
		Address:    address,
		Balance:    usdString,
		Label:      label,
		Type:       badgeType,
		RightColor: color,
	}

	fmt.Println(badge.Coin)

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
