package main

import "fmt"

func main() {
	input := []int{3, 1, 2, 10, 1}
	runningSum(input)

}

func runningSum(nums []int) []int {
	var ans []int
	for i, v := range nums {
		if i == 0 {
			ans = append(ans, v)
			continue
		}
		tmp := ans[i-1] + v
		ans = append(ans, tmp)
	}
	fmt.Println(ans)
	return ans
}
