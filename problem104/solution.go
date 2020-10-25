package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	max := 0
	find(root, 1, &max)
	return max
}

func find(node *TreeNode, current int, max *int) {
	if node == nil {
		return
	}
	if current > *max {
		*max = current
	}
	find(node.Left, current+1, max)
	find(node.Right, current+1, max)
}
