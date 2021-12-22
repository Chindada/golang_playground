package main

import (
	"fmt"
	"math"
)

func main() {
	for {
		fmt.Println(math.Round(0.1*10) / 10)
		fmt.Println(math.Round(0.3*10) / 10)
	}
	// a := []float64{}
	// go func(input *[]float64) {
	// 	for {
	// 		// tmp := []float64{}
	// 		*input = []float64{}
	// 		time.Sleep(time.Second)
	// 		fmt.Println(*input)
	// 	}
	// }(&a)
	// go func(input *[]float64) {
	// 	for {
	// 		// tmp := []float64{}
	// 		*input = []float64{}
	// 		time.Sleep(1 * time.Second)
	// 		fmt.Println(*input)
	// 	}
	// }(&a)
	// go func(input *[]float64) {
	// 	for {
	// 		// tmp := []float64{}
	// 		*input = []float64{}
	// 		time.Sleep(2 * time.Second)
	// 		fmt.Println(*input)
	// 	}
	// }(&a)
	// for {
	// 	a = append(a, 0.1)
	// 	time.Sleep(200 * time.Millisecond)
	// 	fmt.Println(a)
	// }
}
