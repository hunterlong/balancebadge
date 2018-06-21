package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	GetEnv()
	err := LoadEthBlockchains()
	if err != nil {
		panic(err)
	}
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
