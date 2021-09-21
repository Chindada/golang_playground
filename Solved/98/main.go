package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 4,
		},
		Right: &TreeNode{
			Val: 6,
			Left: &TreeNode{
				Val: 3,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	a := isValidBST(&root)
	fmt.Println(a)
}

func isValidBST(root *TreeNode) bool {
	// 中序遍历整个二叉树，如果是搜索树，那么遍历的结果应该是一个升序
	// 中序遍历的顺序是左、中、右
	// 先设置int型的最小值
	pre := ^int(^uint(0) >> 1)
	var midTraverse func(treeNode *TreeNode) bool
	midTraverse = func(treeNode *TreeNode) bool {
		// 当当前结点不为空的时候，此时就有Val值
		if treeNode != nil {
			// 首先遍历左子树
			flag := midTraverse(treeNode.Left)
			if !flag {
				return false
			}
			// 然后遍历中间结点
			// 就算是和前一个结点的值相等也不行，因为这是搜索树，所以数据不能相等
			if treeNode.Val <= pre {
				return false
			}
			// 更新中序遍历的前一个结点
			pre = treeNode.Val
			// 最后遍历右子树
			flag = midTraverse(treeNode.Right)
			if !flag {
				return false
			}
		}
		// 到这个地方就返回true，两种情况可以到这个地方，
		// 1是当当前结点为空
		// 2是当前面的流程走完，也就是遍历完整个二叉树，还没有返回false，那么就返回true
		return true
	}
	return midTraverse(root)
}
