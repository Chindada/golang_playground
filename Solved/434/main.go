package main

import "fmt"

func main() {
	input := "Hello, my name is John"
	countSegments(input)
}

func countSegments(s string) int {
	var ans int
	if len(s) == 0 {
		fmt.Println(ans)
		return ans
	}
	if s[0] != 32 {
		ans = 1
	}
	for i, k := range s {
		if i == len(s)-1 {
			break
		}
		if k == 32 && s[i+1] != 32 {
			ans++
		}
	}
	fmt.Println(ans)
	return ans
}
