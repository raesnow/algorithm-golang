package list

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	newHead := &ListNode{
		Val:  head.Val,
		Next: nil,
	}
	newP := newHead
	for p := head.Next; p != nil; p = p.Next {
		newP.Next = &ListNode{
			Val:  p.Val,
			Next: nil,
		}
		newP = newP.Next
	}
	reversedHead := reverseList(newHead)

	p1, p2 := head, reversedHead
	for ; p1 != nil && p2 != nil; p1, p2 = p1.Next, p2.Next {
		if p1.Val != p2.Val {
			return false
		}
	}
	if p1 != nil || p2 != nil {
		return false
	}
	return true
}

func isPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	slowP, fastP := head, head
	for ; fastP != nil && fastP.Next != nil; slowP, fastP = slowP.Next, fastP.Next.Next {
	}
	if fastP != nil {
		slowP = slowP.Next
	}
	reversedHead := reverseList(slowP)

	p1, p2 := head, reversedHead
	for ; p2 != nil; p1, p2 = p1.Next, p2.Next {
		if p1.Val != p2.Val {
			return false
		}
	}
	return true
}
