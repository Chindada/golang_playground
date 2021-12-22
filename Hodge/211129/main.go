package main

import (
	"211129/restful"
	"fmt"
	"net/http"
	"time"
)

// OrderResponse OrderResponse
type OrderResponse struct {
	Status  string `json:"status"`
	OrderID string `json:"order_id"`
}

// OrderBody OrderBody
type OrderBody struct {
	Stock    string  `json:"stock"`
	Price    float64 `json:"price"`
	Quantity int64   `json:"quantity"`
}

// CancelBody CancelBody
type CancelBody struct {
	OrderID string `json:"order_id"`
}

// OrderAction OrderAction
type OrderAction int64

const (
	// BuyAction BuyAction
	BuyAction OrderAction = iota + 1
	// SellAction SellAction
	SellAction
	// SellFirstAction SellFirstAction
	SellFirstAction
)

type StockOrder struct {
	StockNum  string
	Action    int64
	Quantity  int64
	Price     float64
	OrderID   string
	Status    int64
	TradeTime time.Time
}

func (c *StockOrder) PlaceOrder() {
	var url string
	switch c.Action {
	case int64(BuyAction):
		url = "/pyapi/trade/buy"
	case int64(SellAction):
		url = "/pyapi/trade/sell"
	case int64(SellFirstAction):
		url = "/pyapi/trade/sell_first"
	}
	order := struct {
		Stock    string  `json:"stock"`
		Price    float64 `json:"price"`
		Quantity int64   `json:"quantity"`
	}{
		Stock:    c.StockNum,
		Price:    c.Price,
		Quantity: c.Quantity,
	}
	resp, err := restful.GetClient().R().
		SetBody(order).
		SetResult(&OrderResponse{}).
		Post("http://" + PyServerHost + ":" + PyServerPort + url)
	if err != nil {
		return
	} else if resp.StatusCode() != http.StatusOK {
		return
	}
	res := *resp.Result().(*OrderResponse)
	fmt.Println(res)
}

// SinoPacOrderStatus SinoPacOrderStatus
type SinoPacOrderStatus struct {
	Action    string  `json:"action"`
	Code      string  `json:"code"`
	ID        string  `json:"id"`
	Price     float64 `json:"price"`
	Quantity  int64   `json:"quantity"`
	Status    string  `json:"status"`
	OrderTime string  `json:"order_time"`
}

const (
	// StatusSuccuss StatusSuccuss
	StatusSuccuss string = "success"
	// StatusFail StatusFail
	StatusFail string = "fail"
	// StatusHistoryNotFound StatusHistoryNotFound
	StatusHistoryNotFound string = "history not found"
	// StatusCancelOrderNotFound StatusCancelOrderNotFound
	StatusCancelOrderNotFound string = "cancel order not found"
	// StatusAlreadyCanceled StatusAlreadyCanceled
	StatusAlreadyCanceled string = "order already be canceled"
)

// SinoStatusResponse SinoStatusResponse
type SinoStatusResponse struct {
	Status string               `json:"status"`
	Data   []SinoPacOrderStatus `json:"data"`
}

func (c *StockOrder) FetchOrderStatus() {
	resp, err := restful.GetClient().R().
		SetResult(&SinoStatusResponse{}).
		Get("http://" + PyServerHost + ":" + PyServerPort + "/pyapi/trade/status")
	if err != nil {
		return
	} else if resp.StatusCode() != http.StatusOK {
		return
	}
	res := *resp.Result().(*SinoStatusResponse)
	if res.Status != StatusSuccuss {
		return
	}
}

var (
	PyServerHost string
	PyServerPort string
)

type Sinopac interface {
	PlaceOrder()
	FetchOrderStatus()
}

func main() {
	var qoo Sinopac
	qoo.FetchOrderStatus()
}
