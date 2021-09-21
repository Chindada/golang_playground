package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := ListNode{
		Val: 3,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 0,
				Next: &ListNode{
					Val:  -4,
					Next: &ListNode{},
				},
			},
		},
	}
	detectCycle(&head)
}

func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	tmp := head
	headsMap := make(map[*ListNode]bool)

	for {
		if headsMap[tmp] {
			return tmp
		}
		headsMap[tmp] = true
		tmp = tmp.Next
		if tmp.Next == nil {
			return nil
		}
	}
	return nil
}
