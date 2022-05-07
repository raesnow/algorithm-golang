package list

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	start, end := head, head
	for i := 0; i < k; i++ {
		if end == nil {
			return start
		}
		end = end.Next
	}

	head = reverseGroup(start, end)
	start.Next = reverseKGroup(end, k)
	return head
}

// reverseGroup 反转[head, end)
func reverseGroup(head *ListNode, end *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	if head == end {
		return head
	}

	var pre *ListNode
	p, next := head, head
	for p != end {
		next = p.Next
		p.Next = pre
		pre = p
		p = next
	}
	return pre
}
