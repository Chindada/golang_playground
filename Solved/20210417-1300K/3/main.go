package main

import "fmt"

// You are given a string S consisting of N letters 'a' and/or 'b. In one move, you can swap one letter for
// the other ('a' for 'b' or 'b' for 'a').

// Write a function So lut ion that, given such a string S, returns the minimum number of moves
// required to obtain a string containing no instances of three identical consecutive letters.

// Examples:
// 1. Given S = "baaaaa", the function should return 1. The string without three identical consecutive
// letters which can be obtained in one move is "baabaa".
// 2. Given S = "baaabbaabbba", the function should return 2. There are four valid strings obtainable in
// two moves: for example, "bbaabbaabbaa".
// 3. Given S = "baabab", the function should return 0.

// Write an efficient algorithm for the following assumptions:
//// N is an integer within the range [0..200,000];
//// string S consists only of the characters "a" and/or "b".

func main() {
	S := "baaaaa"
	fmt.Println(Solution(S))
}

func Solution(S string) int {
	ans := 0
	for i := 0; i < len(S); i++ {
		j := i
		for j < len(S) && S[j] == S[i] {
			j++
		}

		ans += (j - i) / 3
		i = j - 1
	}
	return ans
}
