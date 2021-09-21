package main

import (
	"sort"
)

func main() {
	points := [][]int{
		{3, 1},
		{9, 0},
		{1, 0},
		{1, 4},
		{5, 3},
		{8, 8},
	}
	maxWidthOfVerticalArea(points)
}

func maxWidthOfVerticalArea(points [][]int) int {
	var tmp []int
	for _, v := range points {
		tmp = append(tmp, v[0])
	}
	var ans int
	sort.Ints(tmp)
	// tmpAns = tmp[0]
	for i := 0; i < len(tmp)-1; i++ {
		// if tmp[i]-tmpAns > ans {
		// 	ans = tmp[i] - tmpAns
		// }
		// tmpAns = tmp[i]
		aa := tmp[i+1] - tmp[i]
		if aa > ans {
			ans = aa
		}
	}
	return ans
}
