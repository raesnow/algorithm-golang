package list

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	// 检测是否有环
	hasCycle := false
	slowP, fastP := head, head
	for fastP != nil && fastP.Next != nil {
		slowP = slowP.Next
		fastP = fastP.Next.Next
		if slowP == fastP {
			hasCycle = true
			break
		}
	}
	if !hasCycle {
		return nil
	}

	// 寻找环的起始位置
	for slowP = head; slowP != fastP; slowP, fastP = slowP.Next, fastP.Next {

	}
	return slowP
}
