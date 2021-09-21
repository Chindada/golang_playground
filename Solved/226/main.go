package main

// TreeNode TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InvertTree(root *TreeNode) *TreeNode {
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
		rfunc(input.Left)
		rfunc(input.Right)
	}
	rfunc(root)
	return root
}
