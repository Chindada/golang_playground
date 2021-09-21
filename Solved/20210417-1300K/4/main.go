package main

// AND is a standard bitwise operation. For example, given K = 12 (binary representation 01100) and L =
// 21 (binary representation 10101) we obtain:

// 01100 AND
// 10101 =
// 00100

// The AND operation can be extended to N integers, for example:
// 01100 AND
// 10101 AND
// 00100 =
// 00100

// Because the AND product of 01100 (first argument) and 10101 (second argument) is 00100, and the
// AND product of this number with 00100 (third argument) is also 00100.

// Write a function:

////// func Solution(A []int) int

// that, given an array A consisting of N integers, returns the size of the largest possible subset of A
// such that its AND product is greater than 0.

// Examples:
// 1. Given A = [13, 7, 2, 8, 3] your function should return 3.
// We can pick subset 13, 7 and 3. Its AND product is positive and it is the largest possible subset of
// numbers that fulfills the criteria.
// 1101(13) AND
// 0111(7) AND
// 0011(3) =
// 0001(1)

// 2. Given A = [1, 2, 4, 8] your function should return 1. The AND product of each pair from the array is
// equal to 0.

// 3. Given A = [16, 16] your function should return 2. The AND product of both numbers is 16.

// Write an efficient algorithm for the following assumptions:
//// N is an integer within the range [1..100,000);
//// each element of array A is an integer within the range [1..1,000,000,000].

func Solution(A []int) int {
	if len(A) == 0 {
		return 0
	}
	max := A[0]
	for _, val := range A {
		if val > max {
			max = val
		}
	}
	ans := 0
	tmp := 1
	for {
		tmpAns := 0
		for _, v := range A {
			if v&tmp > 0 {
				tmpAns++
			}
		}
		if tmpAns > ans {
			ans = tmpAns
		}
		if max/tmp == 0 {
			break
		}
		tmp *= 2
	}
	return ans
}
