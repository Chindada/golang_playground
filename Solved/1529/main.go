package main

import "fmt"

func main() {
	target := "00000"
	fmt.Println(minFlips(target))
}

func minFlips(target string) int {
	var ans int
	for i, v := range target {
		if ans == 0 {
			if v == 48 {
				continue
			} else {
				ans++
			}
		} else {
			if v != rune(target[i-1]) {
				ans++
			}
		}

	}
	return ans
}
