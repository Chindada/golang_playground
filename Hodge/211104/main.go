package main

import (
	"fmt"
	"time"
)

func main() {
	var a int
	if a = 3; a == 3 {
		fmt.Println("1")
	}
	fmt.Println(a)
	fmt.Println(time.Now().Second())
}
