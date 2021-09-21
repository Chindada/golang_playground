package main

import (
	"sort"
)

func main() {
	piles := []int{9, 8, 7, 6, 5, 1, 2, 3, 4}
	maxCoins(piles)
}

func maxCoins(piles []int) int {
	sort.Ints(piles)
	times := len(piles) / 3
	tmp := 1
	var ans int
	for i := len(piles) - 2; i >= 0; i -= 2 {
		if tmp > times {
			break
		}
		ans += piles[i]
		tmp++
	}
	return ans
}
