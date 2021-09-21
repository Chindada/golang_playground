package main

import "fmt"

// ListNode ListNode
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l1Arr []int
	var l2Arr []int
	tempNodes1 := l1
	tempNodes2 := l2
	for tempNodes1 != nil {
		l1Arr = append(l1Arr, tempNodes1.Val)
		tempNodes1 = tempNodes1.Next
	}
	for tempNodes2 != nil {
		l2Arr = append(l2Arr, tempNodes2.Val)
		tempNodes2 = tempNodes2.Next
	}
	fmt.Println(l1Arr, l2Arr)
	tempArr := compareLength(l1Arr, l2Arr)
	fmt.Println(tempArr)
	headNode := ListNode{
		Val: tempArr[0],
	}
	tail := &headNode
	for i, v := range tempArr {
		if i == 0 {
			continue
		}
		node := ListNode{
			Val: v,
		}
		tail.Next = &node
		tail = &node
	}
	fmt.Println(headNode)
	return &headNode
}

func compareLength(l1Arr, l2Arr []int) (tempArr []int) {
	var spare int
	if len(l1Arr) > len(l2Arr) {
		for i := 0; i < len(l2Arr); i++ {
			tmp := l1Arr[i] + l2Arr[i] + spare
			need := tmp % 10
			spare = tmp / 10
			tempArr = append(tempArr, need)
		}
		for i := len(l2Arr); i < len(l1Arr); i++ {
			tmp := l1Arr[i] + spare
			need := tmp % 10
			spare = tmp / 10
			tempArr = append(tempArr, need)
		}
	} else {
		for i := 0; i < len(l1Arr); i++ {
			tmp := l1Arr[i] + l2Arr[i] + spare
			need := tmp % 10
			spare = tmp / 10
			tempArr = append(tempArr, need)
		}
		for i := len(l1Arr); i < len(l2Arr); i++ {
			tmp := l2Arr[i] + spare
			need := tmp % 10
			spare = tmp / 10
			tempArr = append(tempArr, need)
		}
	}
	if spare != 0 {
		tempArr = append(tempArr, spare)
	}
	return tempArr
}
