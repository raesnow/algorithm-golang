package list

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return nil
	}
	if left == 1 {
		return reverseN(head, right)
	}

	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

var endNodes *ListNode

func reverseN(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if n == 1 {
		endNodes = head.Next
		return head
	}
	last := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = endNodes
	return last
}
