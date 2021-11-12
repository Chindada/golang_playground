package main

func main() {
	root := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 0,
		},
		Right: &TreeNode{
			Val: 0,
		},
	}
	distributeCoins(&root)
}

// TreeNode TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
	var rfunc func(*TreeNode)
	rfunc = func(input *TreeNode) {
		if input == nil {
			return
		}
		if input.Left == nil && input.Right == nil {
			return
		}
		tmpLeft := input.Left
		tmpRight := input.Right
		input.Left = tmpRight
		input.Right = tmpLeft
		if input.Val > 1 {
			if input.Left == nil && input.Right == nil {
				return
			}
		}

		rfunc(input.Left)
		rfunc(input.Right)
	}
	rfunc(root)
	return 0
}
