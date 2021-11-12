package main

import "fmt"

func main() {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	part := 3
	divide := len(input) / part
	fmt.Println(divide)
	var tmp []int
	for i := 1; i <= part; i++ {
		tmp = input[len(input)-i*divide : len(input)-(i-1)*divide]
		fmt.Println(tmp)
	}
}
