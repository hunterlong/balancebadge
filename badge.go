package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	badgeSvg     = string(`<svg xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" width="{{ .Width }}" height="20"> <linearGradient id="b" x2="0" y2="100%"><stop offset="0" stop-color="#bbb" stop-opacity=".1"/> <stop offset="1" stop-opacity=".1"/> </linearGradient> <clipPath id="a"> <rect width="{{ .Width }}" height="20" rx="3" fill="#fff"/> </clipPath> <g clip-path="url(#a)"><path fill="#{{ .LeftColor }}" d="M0 0h{{ .LeftSize }}v20H0z"/> <path fill="#{{.RightColor}}" d="M{{ .LeftSize }} 0h{{ .Width }}v20H{{ .LeftSize }}z"/> <path fill="url(#b)" d="M0 0h{{ .LeftSize }}{{ .Width }}v20H0z"/> </g> <g fill="#fff" text-anchor="middle" font-family="DejaVu Sans,Verdana,Geneva,sans-serif" font-size="110"> <text x="{{ .LeftTextX }}" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="{{ .LeftTextSize }}">{{ .Label }}</text> <text x="{{ .LeftTextX }}" y="140" transform="scale(.1)" textLength="{{ .LeftTextSize }}">{{ .Label }}</text> <text x="{{ .RightTextX }}" y="150" fill="#010101" fill-opacity=".3" transform="scale(.1)" textLength="{{ .RightTextSize }}">{{ .Balance }} {{ .Coin }}</text> <text x="{{ .RightTextX }}" y="140" transform="scale(.1)" textLength="{{ .RightTextSize }}">{{ .Balance }} {{ .Coin }}</text> </g> </svg>`)
	defaultColor = "97CA00"
)

var (
	svgTemplate *template.Template
)

func init() {
	svgTemplate = template.New("badge")
	svgTemplate.Parse(badgeSvg)
}

func NewBadge(r *http.Request) *Badge {
	vars := mux.Vars(r)
	color := r.FormValue("color")
	label := r.FormValue("label")
	badge := &Badge{
		Coin:        strings.ToUpper(vars["coin"]),
		Token:       vars["token"],
		Address:     vars["address"][0:7],
		FullAddress: vars["address"],
		Type:        vars["type"],
		Label:       label,
		Height:      20,
		RightColor:  color,
		LeftColor:   "555555",
		start:       time.Now(),
	}

	badge.Label = badge.Address[0:7]
	if label != "" {
		badge.Label = label
	}

	badge.RightColor = defaultColor
	if color != "" {
		badge.RightColor = ColorToHex(color)
	}
	return badge
}

func (b *Badge) SetBalance() *Badge {
	var balance string
	if b.Coin == "token" {
		balance, _, _ = TokenBalance(b.Coin, b.FullAddress)
	} else {
		balance = CryptoBalance(b.Coin, b.FullAddress)
	}
	b.Balance = balance
	return b
}

func (b *Badge) SetUSD() *Badge {
	usd := FindCoinRate(b.Coin)
	if usd == nil {
		return b
	}
	balFloat, _ := strconv.ParseFloat(b.Balance, 10)
	usdBalance := usd.Price * balFloat
	usdString := fmt.Sprintf("%0.3f", usdBalance)
	b.Balance = usdString
	if b.Token != "" {
		b.Coin = b.Coin + "/USD"
	} else {
		b.Coin = "USD"
	}
	b.toCurrency = true
	return b
}

func (b *Badge) TokenBadge() *Badge {
	balance, symbol, err := TokenBalance(b.Token, b.FullAddress)
	if err != nil {
		fmt.Println(err)
		return b
	}
	b.Balance = balance
	b.Coin = symbol
	return b
}


func (b *Badge) Cache() *Badge {
	switch b.Coin {
	case "BTC":
		b.cache = "public, max-age=900"
	case "LTC":
		b.cache = "public, max-age=600"
	case "ETH":
		b.cache = "public, max-age=120"
	default:
		b.cache = "no-cache, no-store, must-revalidate"
	}
	return b
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
	if b.toCurrency {
		b.Balance = "$" + b.Balance
	}
	b.LeftSize = (len(b.Label) + 1) * 8
	b.LeftTextSize = b.LeftSize * 8
	b.LeftTextX = (b.LeftTextSize / 2) + 60
	b.RightSize = (len(b.Balance) + len(b.Coin) + 1) * 8
	b.RightTextSize = (b.RightSize * 8) + 40
	b.RightTextX = (b.RightTextSize / 2) + (b.LeftSize * 2) + b.LeftTextSize + 60
	b.Width = b.LeftSize + b.RightSize
	return b.Cache()
}
