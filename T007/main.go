package main

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type SinoPacStockVolumeClose struct {
	Code   string  `json:"code"`
	Close  float64 `json:"close"`
	Volume int64   `json:"volume"`
}

type UpdateVolumeArrBody struct {
	StockNumArr []string `json:"stock_num_arr"`
}

type AuthSuces struct {
	ResultA struct {
		Message string `json:"message"`
		Teacher string `json:"teacher"`
		Student string `json:"student"`
		Count   int    `json:"count"`
	} `json:"result_oop"`
}

type AuthSucesB struct {
	ResultCCCC []struct {
		Message string `json:"message"`
		Teacher string `json:"teacher"`
		Student string `json:"student"`
		Count   int    `json:"count"`
	} `json:"result"`
}

type AuthSucesC struct {
	ResultCCCC []struct {
		Code   string  `json:"code"`
		Close  float64 `json:"close"`
		Volume int64   `json:"volume"`
	} `json:"result"`
}

func main() {
	arr := append([]string{}, "233ddssdddddddddddddddd0", "260d9")
	tmp := UpdateVolumeArrBody{
		StockNumArr: arr,
	}

	client := resty.New()
	resp, err := client.R().
		SetBody(&tmp).
		SetResult([]SinoPacStockVolumeClose{}).
		Post("http://127.0.0.1:3333/basic/volumeclose")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Result())
	for _, v := range *resp.Result().(*[]SinoPacStockVolumeClose) {
		fmt.Println(v)
	}
}
