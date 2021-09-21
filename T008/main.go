package main

import (
	"fmt"
	"math"
)

func main() {
	var price float64
	price = 10
	for {
		roundPrice := Round(price, 2)
		if roundPrice > 50 {
			break
		}
		oneTick := GetNewClose(roundPrice, 1)
		fmt.Printf("originadl: %f, one tick: %f, diff: %f, ", roundPrice, oneTick, oneTick-roundPrice)
		fmt.Printf("Buy Price: %d, Sell Price: %d, Diff: %d\n", GetStockBuyCost(roundPrice, 1), GetStockSellCost(oneTick, 1), GetStockSellCost(oneTick, 1)-GetStockBuyCost(roundPrice, 1))
		price += 0.05

	}

}

var TradeQuota int64 = 1000000

var TradeFeeRatio float64 = 0.001425 * 0.38
var TradeTaxRatio float64 = 0.0015

func GetStockBuyCost(price float64, qty int64) int64 {
	return int64(price*float64(qty)*1000 + math.Floor(price*float64(qty)*1000*TradeFeeRatio))
}

func GetStockSellCost(price float64, qty int64) int64 {
	return int64(price*float64(qty)*1000 - math.Floor(price*float64(qty)*1000*TradeFeeRatio) - math.Floor(price*float64(qty)*1000*TradeTaxRatio))
}

func GetNewClose(close float64, unit int64) float64 {
	for {
		var diff float64
		if unit == 0 {
			break
		}
		if close < 10 {
			diff += 0.01
		} else if close < 50 {
			diff += 0.05
		} else if close < 100 {
			diff += 0.1
		} else if close < 500 {
			diff += 0.5
		} else if close < 1000 {
			diff += 1
		} else {
			diff += 5
		}
		if unit > 0 {
			close += diff
			unit--
		} else {
			close -= diff
			unit++
		}
	}
	return Round(close, 2)
}

func Round(val float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(val*p+0.5) / p
}
