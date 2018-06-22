package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"net/http"
	"net/http/httptest"
)

func init() {
	GetEnv()
	err := LoadEthBlockchains()
	if err != nil {
		panic(err)
	}
}

func TestFetchRates(t *testing.T) {
	FetchCoinMarketCap()
}

func TestEthBalance(t *testing.T) {
	balance := EthBalance(eth, "0x004f3e7ffa2f06ea78e14ed2b13e87d710e8013f")
	t.Log("balance: ", balance)
	assert.NotZero(t, balance)
}

func TestBtcBalance(t *testing.T) {
	balance, err := BitcoinBalance(BTCapi, "18CEdeCbtvPivNeUVmwaTMbvCMreFvbe8")
	t.Log("balance: ", balance)
	assert.Nil(t, err)
	assert.NotZero(t, balance)
}

func TestLtcBalance(t *testing.T) {
	balance, err := BitcoinBalance(LTCapi, "LMrS5XQMR233haqDqgs63oAtiWaA2Dj8c9")
	t.Log("balance: ", balance)
	assert.Nil(t, err)
	assert.NotZero(t, balance)
}

func TestEthBadgeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/eth/0x2c70d87ae710b1174ee8c4e2916ea73b789eb37f.svg", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 988, len(rr.Body.Bytes()))
}

func TestEthUsdBadgeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/eth/0x2c70d87ae710b1174ee8c4e2916ea73b789eb37f/usd.svg", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 990, len(rr.Body.Bytes()))
}

func TestBtcBadgeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/btc/141JSRmgzv4km7KT1HdAx5hQnnYJjGb9PG.svg", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 988, len(rr.Body.Bytes()))
}

func TestBtcUsdBadgeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/btc/141JSRmgzv4km7KT1HdAx5hQnnYJjGb9PG/usd.svg", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 990, len(rr.Body.Bytes()))
}

func TestLtcBadgeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/ltc/LLgJTbzZMsRTCUF1NtvvL9SR1a4pVieW89.svg", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 990, len(rr.Body.Bytes()))
}

func TestLtcUsdBadgeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/ltc/LLgJTbzZMsRTCUF1NtvvL9SR1a4pVieW89/usd.svg", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 998, len(rr.Body.Bytes()))
}

func TestTokenBadgeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/token/0xb64ef51c888972c908cfacf59b47c1afbc0ab8ac/0x2c70d87ae710b1174ee8c4e2916ea73b789eb37f.svg", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 992, len(rr.Body.Bytes()))
}

func TestTokenUsdBadgeHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/token/0xb64ef51c888972c908cfacf59b47c1afbc0ab8ac/0x2c70d87ae710b1174ee8c4e2916ea73b789eb37f/usd.svg", nil)
	assert.Nil(t, err)
	rr := httptest.NewRecorder()
	Router().ServeHTTP(rr, req)
	assert.Equal(t, 1004, len(rr.Body.Bytes()))
}