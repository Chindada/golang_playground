package main

// TreeNode TreeNode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func HasPathSum(root *TreeNode, sum int) bool {
	var ans bool
	var rfunc func(*TreeNode, int)
	rfunc = func(input *TreeNode, tempSum int) {
		if input == nil {
			return
		}
		if input.Left == nil && input.Right == nil {
			if input.Val == tempSum {
				ans = true
			}
		}
		if input.Left != nil {
			input.Left.Val += input.Val
			rfunc(input.Left, tempSum)
		}
		if input.Right != nil {
			input.Right.Val += input.Val
			rfunc(input.Right, tempSum)
		}
	}
	rfunc(root, sum)
	return ans
}
