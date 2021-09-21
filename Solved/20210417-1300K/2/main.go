package main

// Write a function so lut ion that, given a string S consisting of N characters, returns the maximum
// number of letters 'a' that can be inserted into S (including at the front and end of S) so that the
// resltin stin de't coxtain three cosecutive letters 'a'. If string s already contains the substring
// "aaa", then your function should return -1.

// Examples:
// 1. Given S = "aabab", the function should return 3, because a string "aabaabaa" can be made.
// 2. Given S = "dog", the function should return 8, because a string "aadaaoaagaa" can be made.
// 3. Given S = "aa", the function should return 0, because no longer string can be made.
// 4. Given S = "baaaa", the function should return -1, because there is a substring "aaa".

// Write an efficient algorithm for the following assumptions:
//// N is an integer within the range [1 ..200,000];
//// string S consists only of lowercase letters (a-z).

func Solution(S string) int {
	ans := 0
	count := 0
	for _, val := range S {
		if val == 97 {
			count++
		} else {
			count = 0
		}
		if count == 3 {
			return -1
		}
	}
	tmp := 0
	for _, val := range S {
		if val != 97 {
			ans += 2 - tmp
			tmp = 0
		} else {
			tmp++
		}
	}
	ans += 2 - tmp
	return ans
}
