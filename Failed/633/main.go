package main

func JudgeSquareSum(c int) bool {
	var a2, b2 []int
	for i := 0; i*i <= c; i++ {
		a2 = append(a2, i)
	}
	b2 = a2
	// fmt.Println(a2)
	// fmt.Println(b2)
	for _, v := range a2 {
		for _, k := range b2 {
			// fmt.Println(v*v, k*k)
			if v*v+k*k == c {
				return true
			}
		}
	}
	return false
}
