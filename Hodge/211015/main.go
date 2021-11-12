package main

import (
	"fmt"
	"time"
)

func main() {
	a := []int{1, 2, 3, 4, 5}
	for {
		a = append(a, a[len(a)-1]+1)
		fmt.Println(a)
		a = a[1:]
		time.Sleep(400 * time.Millisecond)
	}
}
