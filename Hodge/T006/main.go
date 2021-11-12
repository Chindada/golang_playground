package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := struct {
		Response string `json:"response"`
	}{
		Response: "fail",
	}

	aJson, _ := json.Marshal(a)

	fmt.Println(string(aJson))

	res := struct {
		Response string `json:"response"`
	}{}
	json.Unmarshal(aJson, &res)
	fmt.Println(res.Response)
	c := "{\"respddonse\":\"faild\"}"
	fmt.Println(c)
}
