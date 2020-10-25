package main

import "fmt"

func main() {
	root1 := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}
	root2 := &TreeNode{
		Val:  1,
		Left: nil,
		Right: &TreeNode{
			Val:   2,
			Left:  nil,
			Right: nil,
		},
	}
	fmt.Println(isSameTree(root1, root2))
}
