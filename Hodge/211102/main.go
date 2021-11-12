package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		time.Sleep(time.Second)
		a := 2
		b := 3
		if a == 2 {
			fmt.Println("1")
			if b == 3 {
				fmt.Println("2")
			}
			continue
		}
		fmt.Println("3")
	}
}
