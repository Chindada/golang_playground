package main

import "fmt"

func main() {
	edges := [][]int{
		{1, 2},
		{2, 3},
		{4, 2},
	}
	a := findCenter(edges)
	fmt.Println(a)
}

func findCenter(edges [][]int) int {
	a := edges[0][0]
	b := edges[0][1]
	var tmpA, tmpB int
	for _, v := range edges {
		if a == v[0] || a == v[1] {
			tmpA++
		}
		if b == v[0] || b == v[1] {
			tmpB++
		}
	}
	if tmpA > tmpB {
		return a
	} else {
		return b
	}
}
