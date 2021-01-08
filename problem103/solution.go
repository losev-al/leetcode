package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type queueItem struct {
	Node  *TreeNode
	Layer int
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	queue := make([]queueItem, 0)
	queue = append(queue, queueItem{root, 0})
	result := make([][]int, 0)
	for len(queue) > 0 {
		current := queue[0].Node
		layer := queue[0].Layer
		queue = queue[1:]
		if current != nil {
			if len(result) == layer {
				result = append(result, make([]int, 0))
			}
			if layer%2 == 0 {
				result[layer] = append(result[layer], current.Val)
			} else {
				result[layer] = append([]int{current.Val}, result[layer]...)
			}
			queue = append(queue, queueItem{current.Left, layer + 1})
			queue = append(queue, queueItem{current.Right, layer + 1})
		}
	}
	return result
}
