package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEthBalance(t *testing.T) {
	balance := EthBalance("0x4f70Dc5Da5aCf5e71905c3a8473a6D8a7E7Ba4c5")
	assert.NotZero(t, balance)
}
