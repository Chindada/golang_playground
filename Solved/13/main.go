package main

import "fmt"

func main() {
	s := "MCMXCIV"
	ans := romanToInt(s)
	fmt.Println(ans)
}

func romanToInt(s string) int {

	num := make(map[byte]int)
	num['M'] = 1000
	num['D'] = 500
	num['C'] = 100
	num['L'] = 50
	num['X'] = 10
	num['V'] = 5
	num['I'] = 1

	stringLen := len(s)

	if stringLen == 0 {
		return 0
	}

	ans := num[s[stringLen-1]]
	for i := stringLen - 2; i >= 0; i-- {
		n := num[s[i]]

		if n < num[s[i+1]] {
			ans -= n
		} else {
			ans += n
		}
	}

	return ans
}
