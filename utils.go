package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"time"
)

func GetEnv() {
	BTCapi = os.Getenv("BTC")
	BTCTESTapi = os.Getenv("BTCTEST")
	LTCapi = os.Getenv("LTC")
	LTCTESTapi = os.Getenv("LTCTEST")
	ETHapi = os.Getenv("ETH")
	ROPSTENapi = os.Getenv("ROPSTEN")
	RINKEBYapi = os.Getenv("RINKEBY")
	DASHapi = os.Getenv("DASH")
	ZCASHapi = os.Getenv("ZCASH")
	BCHapi = os.Getenv("BCH")
	BCHTESTapi = os.Getenv("BCHTEST")
}

func CoinMarketCapTicker() {
	defer CoinMarketCapTicker()
	fmt.Println("Loading Coin Market Cap Rates")
	FetchCoinMarketCap()
	time.Sleep(5 * time.Minute)
}

var renderFloatPrecisionMultipliers = [10]float64{
	1,
	10,
	100,
	1000,
	10000,
	100000,
	1000000,
	10000000,
	100000000,
	1000000000,
}

var renderFloatPrecisionRounders = [10]float64{
	0.5,
	0.05,
	0.005,
	0.0005,
	0.00005,
	0.000005,
	0.0000005,
	0.00000005,
	0.000000005,
	0.0000000005,
}

func RenderFloat(format string, amount string) string {
	n, _ := strconv.ParseFloat(amount, 10)

	if math.IsNaN(n) {
		return "NaN"
	}
	if n > math.MaxFloat64 {
		return "Infinity"
	}
	if n < -math.MaxFloat64 {
		return "-Infinity"
	}

	precision := 2
	decimalStr := "."
	thousandStr := ","
	positiveStr := ""
	negativeStr := "-"

	if len(format) > 0 {
		precision = 9
		thousandStr = ""
		formatDirectiveChars := []rune(format)
		formatDirectiveIndices := make([]int, 0)
		for i, char := range formatDirectiveChars {
			if char != '#' && char != '0' {
				formatDirectiveIndices = append(formatDirectiveIndices, i)
			}
		}

		if len(formatDirectiveIndices) > 0 {
			if formatDirectiveIndices[0] == 0 {
				if formatDirectiveChars[formatDirectiveIndices[0]] != '+' {
					panic("RenderFloat(): invalid positive sign directive")
				}
				positiveStr = "+"
				formatDirectiveIndices = formatDirectiveIndices[1:]
			}
			if len(formatDirectiveIndices) == 2 {
				if (formatDirectiveIndices[1] - formatDirectiveIndices[0]) != 4 {
					panic("RenderFloat(): thousands separator directive must be followed by 3 digit-specifiers")
				}
				thousandStr = string(formatDirectiveChars[formatDirectiveIndices[0]])
				formatDirectiveIndices = formatDirectiveIndices[1:]
			}
			if len(formatDirectiveIndices) == 1 {
				decimalStr = string(formatDirectiveChars[formatDirectiveIndices[0]])
				precision = len(formatDirectiveChars) - formatDirectiveIndices[0] - 1
			}
		}
	}

	var signStr string
	if n >= 0.000000001 {
		signStr = positiveStr
	} else if n <= -0.000000001 {
		signStr = negativeStr
		n = -n
	} else {
		signStr = ""
		n = 0.0
	}

	intf, fracf := math.Modf(n + renderFloatPrecisionRounders[precision])
	intStr := strconv.Itoa(int(intf))
	if len(thousandStr) > 0 {
		for i := len(intStr); i > 3; {
			i -= 3
			intStr = intStr[:i] + thousandStr + intStr[i:]
		}
	}
	if precision == 0 {
		return signStr + intStr
	}
	fracStr := strconv.Itoa(int(fracf * renderFloatPrecisionMultipliers[precision]))
	if len(fracStr) < precision {
		fracStr = "000000000000000"[:precision-len(fracStr)] + fracStr
	}
	return signStr + intStr + decimalStr + fracStr
}
