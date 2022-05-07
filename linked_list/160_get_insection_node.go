package list

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	if headA == headB {
		return headA
	}

	for pA, pB := headA, headB; pA != pB; {
		pA = pA.Next
		pB = pB.Next
		if pA == nil && pB == nil {
			return nil
		}
		if pA == nil {
			pA = headB
		}
		if pB == nil {
			pB = headA
		}
		if pA == pB {
			return pA
		}
	}
	return nil
}
