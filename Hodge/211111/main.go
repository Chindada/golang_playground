package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var arr []float64
	var max, min float64
	for {
		var result int
		arr = append(arr, rand.Float64()*3+50)
		if len(arr) > 11 {
			arr = arr[1:]
		} else if len(arr) < 11 {
			continue
		}
		var totalRsi float64
		for _, v := range arr {
			totalRsi += v
		}
		average := totalRsi / float64(len(arr))
		for _, v := range arr {
			if v > average {
				result++
			}
		}
		tmp := float64(result) / float64(len(arr))
		if max == 0 && min == 0 {
			max = tmp
			min = tmp
		}
		// if tmp > max {
		// 	max = tmp
		// 	fmt.Println(tmp)
		// 	// fmt.Println(arr)
		// }
		// if tmp < min {
		// 	min = tmp
		// 	fmt.Println(tmp)
		// 	// fmt.Println(arr)
		// }
		fmt.Println(tmp)
	}
}
