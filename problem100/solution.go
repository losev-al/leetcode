package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	chanA := make(chan *TreeNode)
	chanB := make(chan *TreeNode)
	go fill(p, chanA, true)
	go fill(q, chanB, true)
	for firstNode := range chanA {
		if secondNode, ok := <-chanB; ok {
			if firstNode == nil && secondNode != nil ||
				firstNode != nil && secondNode == nil ||
				firstNode != nil && secondNode != nil && firstNode.Val != secondNode.Val {
				return false
			}
		} else {
			return false
		}
	}
	if _, ok := <-chanB; ok {
		return false
	}
	return true
}

func fill(p *TreeNode, a chan *TreeNode, isRoot bool) {
	a <- p
	if p != nil {
		fill(p.Left, a, false)
		fill(p.Right, a, false)
	}
	if isRoot {
		close(a)
	}
}
