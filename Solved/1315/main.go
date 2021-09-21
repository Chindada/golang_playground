package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SumEvenGrandparent(root *TreeNode) int {
	var nodesSum []int
	var ans int
	var rfunc func(*TreeNode)
	rfunc = func(input *TreeNode) {
		if input == nil {
			return
		}
		if input.Val%2 == 0 {
			var tmpSum int
			if input.Left != nil {
				if input.Left.Left != nil {
					tmpSum += input.Left.Left.Val
				}
				if input.Left.Right != nil {
					tmpSum += input.Left.Right.Val
				}
			}
			if input.Right != nil {
				if input.Right.Left != nil {
					tmpSum += input.Right.Left.Val
				}
				if input.Right.Right != nil {
					tmpSum += input.Right.Right.Val
				}
			}
			nodesSum = append(nodesSum, tmpSum)
		}
		rfunc(input.Left)
		rfunc(input.Right)
	}
	rfunc(root)
	for _, v := range nodesSum {
		ans += v
	}
	return ans
}

func SumEvenGrandparent2(root *TreeNode) int {
	ans := 0
	Helper(root, &ans, 0, nil, nil)
	return ans
}

// 98.84% faster and 66.28% memory.
func Helper(root *TreeNode, ans *int, depth int, parent, grandpa *TreeNode) {
	if depth >= 2 {
		if grandpa != nil && grandpa.Val%2 == 0 {
			*ans += root.Val
		}
	}
	if root.Left != nil {
		Helper(root.Left, ans, depth+1, root, parent)
	}
	if root.Right != nil {
		Helper(root.Right, ans, depth+1, root, parent)
	}
}
