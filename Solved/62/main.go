package main

import "fmt"

func main() {
	m := 23
	n := 12
	fmt.Println(uniquePaths(m, n))

}

func uniquePaths(m int, n int) int {
	if m == 1 || n == 1 {
		return 1
	}
	total := m - 1 + n - 1
	var ans, b, c int
	if m > n {
		b = m - 1
		c = n - 1
	} else {
		b = n - 1
		c = m - 1
	}
	ans = 1
	for {
		if total == 1 {
			break
		}
		if total > b {
			ans *= total
		}
		if total <= c {
			ans /= total
		}
		total--
	}
	return ans
}

func UniquePaths2(m int, n int) int {
	if m < 1 || n < 1 {
		return 0
	}
	table := make([]int, n)
	for j := 0; j < n; j++ {
		table[j] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			table[j] += table[j-1]
		}
	}
	return table[n-1]
}
