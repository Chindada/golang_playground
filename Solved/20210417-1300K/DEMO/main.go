package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{1, 2, 3}
	b := Solution(a)
	fmt.Println(b)
}

func Solution(A []int) int {
	// write your code in Go 1.4
	sort.Ints(A)
	fmt.Println(A)
	for i, v := range A {
		if i == 0 {
			continue
		}
		if v == A[i-1] {
			continue
		}
		if A[0] < 0 {
			return 1
		}
		if v-A[i-1] > 1 {
			return A[i-1] + 1
		}
	}
	return A[len(A)-1] + 1
}
