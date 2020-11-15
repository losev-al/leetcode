package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	path1 := findPath(root, p)
	path2 := findPath(root, q)
	for _, node1 := range path1 {
		for _, node2 := range path2 {
			if node1 == node2 {
				return node1
			}
		}
	}
	return nil
}

func findPath(node, targetNode *TreeNode) (path []*TreeNode) {
	if node == targetNode {
		path = append(path, node)
		return
	}
	if node.Left != nil {
		if path = findPath(node.Left, targetNode); path != nil {
			path = append(path, node)
			return
		}
	}
	if node.Right != nil {
		if path = findPath(node.Right, targetNode); path != nil {
			path = append(path, node)
			return
		}
	}
	return nil
}
