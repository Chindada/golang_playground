package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	dateTime := "2021-09-13 02:13:20.123456"
	timeStamp, _ := MicroDateTimeToTimeStamp(dateTime)
	tradeTime := time.Unix(0, timeStamp)
	fmt.Println(tradeTime.UTC())
	fmt.Println(tradeTime.Local())
	fmt.Println(tradeTime.After(time.Now()))
}

const (
	// TradeYear TradeYear
	TradeYear int64 = 2021
	// OneTimeQuantity OneTimeQuantity
	OneTimeQuantity int64 = 1
	// LongTimeLayout LongTimeLayout
	LongTimeLayout string = "2006-01-02 15:04:05"
	// ShortTimeLayout ShortTimeLayout
	ShortTimeLayout string = "2006-01-02"
)

// MicroDateTimeToTimeStamp MicroDateTimeToTimeStamp
func MicroDateTimeToTimeStamp(dateTime string) (timeStamp int64, err error) {
	dataTimeStamp, err := time.ParseInLocation(LongTimeLayout, dateTime[:19], time.Local)
	if err != nil {
		return timeStamp, err
	}
	microSec, err := StrToInt64(dateTime[20:])
	if err != nil {
		return timeStamp, err
	}
	return dataTimeStamp.UnixNano() + microSec*1000, err
}

// StrToInt64 StrToInt64
func StrToInt64(input string) (ans int64, err error) {
	ans, err = strconv.ParseInt(input, 10, 64)
	if err != nil {
		return ans, err
	}
	return ans, err
}
