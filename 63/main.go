package main

import "fmt"

func main() {
	obstacleGrid := [][]int{
		{0, 0, 0, 0},
		{0, 1, 0, 0},
		{0, 0, 0, 0},
		{0, 0, 1, 0},
		{0, 0, 0, 0},
	}
	fmt.Println("ANSSSS", uniquePathsWithObstacles(obstacleGrid))
}

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	total := uniquePaths(m, n)
	var obs [][]int
	for a, v := range obstacleGrid {
		for b, k := range v {
			if k == 1 {
				tmp := []int{a + 1, b + 1}
				obs = append(obs, tmp)
			}
		}
	}
	var obsPath int
	for _, v := range obs {
		part1 := uniquePaths(v[0], v[1])
		part2 := uniquePaths(m-v[0]+1, n-v[1]+1)
		obsPath += (part1 * part2)
	}
	fmt.Println(total, obsPath)
	if total-obsPath < 0 {
		return 0
	}
	return total - obsPath
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
