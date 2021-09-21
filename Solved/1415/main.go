package main

import "fmt"

func main() {
	n := 1
	k := 2
	// for k := 1; k < 10; k++ {
	// 	fmt.Println(getHappyString(n, k))
	// }
	// getHappyString(n, k)
	// fmt.Println(getHappyString(n, k))
	// var ttt []byte
	// ttt = append(ttt, 97, 98, 99)
	// fmt.Println(string(ttt))
	fmt.Println(getHappyString(n, k))
}

func getHappyString(n int, k int) string {
	if n == 1 {
		if k == 1 {
			return "a"
		}
		if k == 2 {
			return "b"
		}
		if k == 3 {
			return "c"
		}
		if k > 3 {
			return ""
		}
	}
	branchNumber := 1
	for i := 1; i <= n-1; i++ {
		branchNumber *= 2
	}
	if k > branchNumber*3 {
		return ""
	}
	var position int
	var start string
	if k <= branchNumber {
		start = "a"
		position = k - 1
	} else if k > branchNumber && k <= branchNumber*2 {
		start = "b"
		position = k - branchNumber - 1
	} else {
		start = "c"
		position = k - branchNumber*2 - 1
	}
	var tmp []string
	tmp = append(tmp, start)
	// tmp = append(tmp, "b")
	// tmp = append(tmp, "c")
	var letters []string
	for {
		letters = nil
		for _, v := range tmp {
			if v[len(v)-1:] == "a" {
				tt := v + "b"
				yy := v + "c"
				letters = append(letters, tt)
				letters = append(letters, yy)
			}
			if v[len(v)-1:] == "b" {
				tt := v + "a"
				yy := v + "c"
				letters = append(letters, tt)
				letters = append(letters, yy)
			}
			if v[len(v)-1:] == "c" {
				tt := v + "a"
				yy := v + "b"
				letters = append(letters, tt)
				letters = append(letters, yy)
			}
		}
		tmp = letters
		if len(letters) == branchNumber {
			break
		}
	}
	return letters[position]
}

func GetHappyString2(n int, k int) string {
	tot := 3
	for i := 1; i < n; i++ {
		tot <<= 1
	}
	if tot < k {
		return ""
	}
	ans := make([]byte, n)
	k--
	ans[0] = byte('a' + k/(tot/3))
	tot /= 3
	k %= tot
	for i := 1; i < n; i++ {
		ans[i] = 'a'
		if ans[i-1] == ans[i] {
			ans[i]++
		}
		tot /= 2
		if k >= tot {
			ans[i]++
		}
		if ans[i-1] == ans[i] {
			ans[i]++
		}
		k %= tot
	}
	return string(ans)
}

func GetHappyString3(n int, k int) string {
	if n == 1 {
		if k == 1 {
			return "a"
		}
		if k == 2 {
			return "b"
		}
		if k == 3 {
			return "c"
		}
		if k > 3 {
			return ""
		}
	}
	branchNumber := 1
	for i := 1; i <= n-1; i++ {
		branchNumber *= 2
	}
	if k > branchNumber*3 {
		return ""
	}
	var letters [][]string
	letters = append(letters, []string{"a"})
	letters = append(letters, []string{"b"})
	letters = append(letters, []string{"c"})
	for {
		var tmp [][]string
		for i := 0; i < len(letters); i++ {
			aa := make([]string, len(letters[i]))
			bb := make([]string, len(letters[i]))
			copy(aa, letters[i])
			copy(bb, letters[i])
			if letters[i][len(letters[i])-1] == "a" {
				aa = append(aa, "b")
				bb = append(bb, "c")
			} else if letters[i][len(letters[i])-1] == "b" {
				aa = append(aa, "a")
				bb = append(bb, "c")
			} else if letters[i][len(letters[i])-1] == "c" {
				aa = append(aa, "a")
				bb = append(bb, "b")
			}
			tmp = append(tmp, aa)
			tmp = append(tmp, bb)
		}
		letters = tmp
		if len(letters) == branchNumber*3 {
			break
		}
	}
	return "ans"
}
