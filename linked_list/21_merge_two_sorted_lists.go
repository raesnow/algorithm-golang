package list

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	dummy := &ListNode{}
	p := dummy
	p1, p2 := list1, list2
	for ; p1 != nil && p2 != nil; p = p.Next {
		if p1.Val <= p2.Val {
			p.Next = p1
			p1 = p1.Next
		} else {
			p.Next = p2
			p2 = p2.Next
		}
	}

	for ; p1 != nil; p1, p = p1.Next, p.Next {
		p.Next = p1
	}
	for ; p2 != nil; p2, p = p2.Next, p.Next {
		p.Next = p2
	}
	return dummy.Next
}
