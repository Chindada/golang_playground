package main

import "fmt"

func main() {
	grid := [][]int{
		{3, 0, 8, 4},
		{2, 4, 5, 7},
		{9, 2, 6, 3},
		{0, 3, 1, 0},
	}
	maxIncreaseKeepingSkyline(grid)
}

func maxIncreaseKeepingSkyline(grid [][]int) int {
	var ans int
	top := make([]int, len(grid[0]))
	left := make([]int, len(grid))
	for i, v := range grid {
		for j, k := range v {
			if k > top[j] {
				top[j] = k
			}
			if k > left[i] {
				left[i] = k
			}
		}
	}
	fmt.Println(top, left)
	for i, v := range grid {
		for j, k := range v {
			if top[j] < left[i] {
				if top[j]-k > 0 {
					ans += (top[j] - k)
				}
			} else {
				if left[i]-k > 0 {
					ans += (left[i] - k)
				}
			}
		}
	}
	fmt.Println(ans)
	return ans
}
