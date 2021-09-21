package main

import (
	"fmt"
	"strconv"
)

func main() {
	ids := []int{1, 2, 3}
	headers := make(map[string]string)
	var header string
	for i, v := range ids {
		text := strconv.Itoa(v)
		if i == len(ids)-1 {
			header += text
		} else {
			header += text
			header += ","
		}
	}
	headers["ids"] = "[" + header + "]"
	fmt.Println(headers["ids"])
}
