package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type nodeInfo struct {
	Val     int
	isRight bool
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	path := make([]nodeInfo, 0, 100)

	return check(root.Left, append(path, nodeInfo{root.Val, false})) &&
		check(root.Right, append(path, nodeInfo{root.Val, true}))
}

func check(root *TreeNode, path []nodeInfo) bool {
	if root == nil {
		return true
	}
	return checkPath(root.Val, path) &&
		check(root.Left, append(path, nodeInfo{root.Val, false})) &&
		check(root.Right, append(path, nodeInfo{root.Val, true}))
}

func checkPath(val int, infos []nodeInfo) bool {
	result := true
	for _, info := range infos {
		result = result && (info.Val < val == info.isRight) && info.Val != val
	}
	return result
}
