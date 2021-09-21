package main

import (
	"fmt"
	"regexp"
	"strings"
	"time"
)

func main() {
	// var date string
	// nowString := time.Now().String()
	// tm, _ := time.ParseInLocation("2006-01-02 15:04:05", nowString[:10]+" 13:30:00", time.Local)
	// if time.Now().Before(tm) {
	// 	date = time.Now().AddDate(0, 0, 1).String()[:10]
	// } else {
	// 	date = nowString[:10]
	// }
	s := "[Trade(contract=Contract(security_type=<SecurityType.Stock: 'STK'>, exchange=<Exchange.TSE: 'TSE'>, code='2890'), order=Order(action=<Action.Buy: 'Buy'>, price=11.0, quantity=1, id='888c19ad', seqno='144802', ordno='     ', account=Account(account_type=<AccountType.Stock: 'S'>, person_id='F127522501', broker_id='9A89', account_id='2263265', signed=True), price_type=<StockPriceType.LMT: 'LMT'>, order_type=<FuturesOrderType.ROD: 'ROD'>), status=OrderStatus(id='888c19ad', status=<Status.Cancelled: 'Cancelled'>, status_code='X', order_datetime=datetime.datetime(2021, 6, 18, 22, 43, 59), cancel_quantity=1, deals=[])), Trade(contract=Contract(security_type=<SecurityType.Stock: 'STK'>, exchange=<Exchange.TSE: 'TSE'>, code='2890'), order=Order(action=<Action.Buy: 'Buy'>, price=11.0, quantity=1, id='9367ccec', seqno='146448', ordno='     ', account=Account(account_type=<AccountType.Stock: 'S'>, person_id='F127522501', broker_id='9A89', account_id='2263265', signed=True), price_type=<StockPriceType.LMT: 'LMT'>, order_type=<FuturesOrderType.ROD: 'ROD'>), status=OrderStatus(id='9367ccec', status=<Status.PreSubmitted: 'PreSubmitted'>, status_code='R', order_datetime=datetime.datetime(2021, 6, 19, 16, 6, 59), deals=[])), Trade(contract=Contract(security_type=<SecurityType.Stock: 'STK'>, exchange=<Exchange.TSE: 'TSE'>, code='2890'), order=Order(action=<Action.Buy: 'Buy'>, price=11.0, quantity=1, id='ef7fd535', seqno='146432', ordno='     ', account=Account(account_type=<AccountType.Stock: 'S'>, person_id='F127522501', broker_id='9A89', account_id='2263265', signed=True), price_type=<StockPriceType.LMT: 'LMT'>, order_type=<FuturesOrderType.ROD: 'ROD'>), status=OrderStatus(id='ef7fd535', status=<Status.PreSubmitted: 'PreSubmitted'>, status_code='R', order_datetime=datetime.datetime(2021, 6, 19, 16, 2, 48), deals=[]))]"
	a := regexp.MustCompile(`Trade`)
	reg := a.Split(s, -1)
	for _, v := range reg {
		tmp := strings.Replace(v, ",", "", -1)
		tmp = strings.Replace(tmp, "'", "", -1)
		tmp = strings.Replace(tmp, "<", "", -1)
		tmp = strings.Replace(tmp, ">", "", -1)
		tmp = strings.Replace(tmp, "(", "", -1)
		tmp = strings.Replace(tmp, ")", "", -1)
		fields := strings.Fields(tmp)
		if len(fields) < 33 {
			continue
		}
		stockNum := fields[4][5:]
		action := fields[6]
		price := fields[7][6:]
		status := fields[24]
		statusID := fields[22][21:]
		fmt.Println(stockNum, action, price, status, statusID)
	}
}

func IsTimeThrough(clockTime string) bool {
	tm, err := time.Parse("2006-01-02 15:04:05", clockTime)
	if err != nil {
		return false
	}

	if time.Now().Add(time.Minute * 10).Before(tm) {
		return true
	}

	return false
}
