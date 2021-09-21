package main

import "fmt"

func LengthOfLongestSubstring(s string) int {
	var a int
	var tmp int
	lenthOfString := len(s)
	uniqueCharMap := make(map[byte]bool)
	for i := 0; i < lenthOfString; i++ {
		if !uniqueCharMap[s[i]] {
			uniqueCharMap[s[i]] = true
			tmp++
		} else {
			if tmp > a {
				a = tmp
			}
			i -= (tmp - 1)
			uniqueCharMap = make(map[byte]bool)
			uniqueCharMap[s[i]] = true
			tmp = 1
		}
	}
	if tmp > a {
		a = tmp
	}
	fmt.Println("Answer", a)
	return a
}
