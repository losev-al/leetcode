package main

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	i := 0
	temp := head
	separator := head
	for temp.Next != nil {
		if i >= k {
			separator = separator.Next
		}
		i++
		temp = temp.Next
	}
	temp.Next = head
	if i < k {
		k = k % (i + 1)
		i -= k
		for i > 0 {
			i--
			separator = separator.Next
		}
	}
	result := separator.Next
	separator.Next = nil
	return result
}
