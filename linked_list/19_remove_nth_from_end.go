package list

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == nil {
		return nil
	}

	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	slowP, fastP := dummy, dummy
	for i := 0; i < n; i++ {
		fastP = fastP.Next
		if fastP == nil {
			return head
		}
	}

	for ; fastP.Next != nil; slowP, fastP = slowP.Next, fastP.Next {

	}
	slowP.Next = slowP.Next.Next
	return dummy.Next
}
