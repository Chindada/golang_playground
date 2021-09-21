package main

import "fmt"

func main() {
	input := ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}
	middleNode(&input)
}

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}
	var i, j int
	tmp := head
	i = 1
	for {
		tmp = tmp.Next
		i++
		if tmp.Next == nil {
			break
		}
	}
	j = 1
	for {
		head = head.Next
		j++
		if j == (i/2)+1 {
			fmt.Println(head.Val)
			return head
		}
	}
}
