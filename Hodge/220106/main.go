package main

import (
	"fmt"
	"sort"
)

func main() {
	// a := time.Time{}
	// b := time.Now()
	// fmt.Println(a, b)
	// fmt.Println(b.Equal(a))
	a := []int64{1, 2, 3, 4, 5, 6}
	fmt.Println(a[3:])
	// b := GetPRByVolume(5, a)
	// fmt.Println(b)
}

// GetPRByVolume GetPRByVolume
func GetPRByVolume(volume int64, c []int64) float64 {
	if len(c) < 2 {
		return 0
	}

	sort.Slice(c, func(i, j int) bool {
		return c[i] > c[j]
	})
	fmt.Println(c)
	var position int
	for i, v := range c {
		if volume >= v {
			position = i
			fmt.Println(v)
			break
		}
		if i == len(c)-1 && position == 0 {
			position = len(c)
		}
	}
	fmt.Println(position)
	return 100 * float64(len(c)-position) / float64(len(c))
}
