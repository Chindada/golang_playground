package main

import (
	"fmt"
)

func main() {
	input := []int{2, 3, 6, 7}
	fmt.Println(combinationSum(input, 7))
}

func combinationSum(candidates []int, target int) [][]int {
	ans := [][]int{}
	for i, v := range candidates {
		l := 1
		for {
			if v*l > target {
				break
			}
			if target%v*l == 0 {
				times := target / (v * l)
				var tmp []int
				for i := 0; i < l; i++ {
					tmp = append(tmp, v)
				}
				for i := 0; i < times; i++ {
					tmp = append(tmp, v)
				}
				ans = append(ans, tmp)
			} else {
				for j, k := range candidates {
					if j == i {
						continue
					}
					if target%v == k {
						var tmp []int
						times := target / v
						for i := 0; i < times; i++ {
							tmp = append(tmp, v)
						}
						tmp = append(tmp, k)
						ans = append(ans, tmp)
					}
				}
			}
		}
	}

	return ans
}
