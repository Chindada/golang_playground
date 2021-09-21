package main

func TwoSum(nums []int, target int) []int {
	output := make([]int, 2)
	for i, num := range nums {
		for k, secondNum := range nums {
			if k == i {
				continue
			}
			sum := num + secondNum
			if sum == target {
				output[0] = i
				output[1] = k
			}
		}
	}
	return output
}
