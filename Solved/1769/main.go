package main

import "fmt"

func main() {
	boxes := "001011"
	minOperations(boxes)
}

func minOperations(boxes string) []int {
	ans := make([]int, len(boxes))
	for i, v := range boxes {
		fmt.Println(v)
		for j, k := range boxes {
			if j == i || k == 48 {
				continue
			}
			if j > i {
				ans[i] += (j - i)
			} else {
				ans[i] += (i - j)
			}
		}
	}
	fmt.Println(ans)
	return ans
}
