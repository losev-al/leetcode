package main

import "fmt"

func main() {
	node5 := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val:   6,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}

	node1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   0,
			Left:  nil,
			Right: nil,
		},
		Right: &TreeNode{
			Val:   8,
			Left:  nil,
			Right: nil,
		},
	}

	root := &TreeNode{
		Val:   3,
		Left:  node5,
		Right: node1,
	}
	fmt.Println(lowestCommonAncestor(root, node5, node1))
}
