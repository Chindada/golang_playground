package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	tmp := head
	headsMap := make(map[*ListNode]bool)

	for {
		if headsMap[tmp] {
			return true
		}
		headsMap[tmp] = true
		tmp = tmp.Next
		if tmp.Next == nil {
			return false
		}
	}
}
