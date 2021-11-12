package main

import (
	"fmt"
	"math"
)

var TradeFeeRatio float64 = 0.001425
var TradeTaxRatio float64 = 0.0015

func main() {
	// for i := 1; i < 50; i++ {
	// fmt.Println("---")
	// fmt.Println(GetNewClose(float64(i), 2), "000")
	a := GetNewClose(23.2, -3)
	fmt.Println(a, Round(a, 2))
	// }
	// fmt.Println(GetStockBuyCost(26.45, 1))
	// fmt.Println(GetStockSellCost(26.3, 1))
}

func GetStockBuyCost(price float64, qty int64) int64 {
	return int64(price*float64(qty)*1000 + math.Floor(price*float64(qty)*1000*TradeFeeRatio))
}

func GetStockSellCost(price float64, qty int64) int64 {
	return int64(price*float64(qty)*1000 - math.Floor(price*float64(qty)*1000*TradeFeeRatio) - math.Floor(price*float64(qty)*1000*TradeTaxRatio))
}

func Round(val float64, precision int) float64 {
	p := math.Pow10(precision)
	return math.Floor(val*p+0.5) / p
}

func GetNewClose(close float64, unit int64) float64 {
	for {
		var diff float64
		if unit == 0 {
			break
		}
		if close < 10 {
			diff = 0.01
		} else if close < 50 {
			diff = 0.05
		} else if close < 100 {
			diff = 0.1
		} else if close < 500 {
			diff = 0.5
		} else if close < 1000 {
			diff = 1
		} else {
			diff = 5
		}
		if unit > 0 {
			close += diff
			unit--
		} else {
			close -= diff
			unit++
		}
	}
	// return Round(close, 2)
	return close
}
