package main

import "fmt"

func main() {
	a := []int{1, 3, 5, 6}
	b := 0
	searchInsert(a, b)
}

func searchInsert(nums []int, target int) int {
	var ans int
	for i, v := range nums {
		if i == 0 && target < v {
			ans = 0
			break
		}
		if target == v {
			ans = i
			break
		}
		if i == len(nums)-1 {
			ans = i + 1
		} else if target > v && target < nums[i+1] {
			ans = i + 1
			break
		}
	}
	fmt.Println(ans)
	return ans
}
