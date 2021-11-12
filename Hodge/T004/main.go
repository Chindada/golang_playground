package main

import (
	"errors"
	"fmt"
)

func main() {
	a := map[int]string{
		1: "2313",
		2: "12313",
		3: "asdfa",
	}
	for key := range a {
		fmt.Println(key, a)
	}
	panic(errors.New("123"))
}
