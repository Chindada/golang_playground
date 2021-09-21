package main

import "fmt"

func Rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] < nums[1] {
			return nums[1]
		}
		return nums[0]
	}
	var totalMoney int
	var totalMoney2 int
	var i int
	for {
		if i > len(nums)-1 {
			break
		}
		totalMoney += nums[i]
		fmt.Println("++++", nums[i], i)
		j := i + 2
		// var tmp int
		for {
			if j+1 > len(nums)-1 {
				break
			}
			// fmt.Println(j)
			// tmp += nums[j]
			// if tmp > nums[j+1] {
			// 	j++
			// 	break
			// }
			if nums[j+1] > nums[j] {
				j++
			} else {
				break
			}
		}
		i = j
	}
	i = 1
	for {
		if i > len(nums)-1 {
			break
		}
		totalMoney2 += nums[i]
		fmt.Println("+++++++++", nums[i], i)
		j := i + 2
		var tmp int
		for {
			if j+1 > len(nums)-1 {
				break
			}
			// fmt.Println(j)
			tmp += nums[j]
			if tmp > nums[j+1] {
				break
			}
			j++
		}
		i = j
	}
	fmt.Println(totalMoney)
	fmt.Println(totalMoney2)
	if totalMoney2 > totalMoney {
		return totalMoney2
	}
	return totalMoney
}
