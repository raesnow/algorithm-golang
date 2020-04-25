package main

/*
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
 */

type ListNode struct {
	Val int
	Next *ListNode
}

// addTwoNumbers
// time: O(max(m,n))
// storage: O(max(m,n))
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	l3Head := &ListNode{}
	l3 := l3Head
	carry := 0
	for ; l1 != nil && l2 != nil; l1, l2, l3 = l1.Next, l2.Next, l3.Next {
		sum := l1.Val + l2.Val + carry
		l3.Next = &ListNode{
			Val:  sum % 10,
		}
		carry = sum / 10
	}

	if l2 != nil {
		l1 = l2
	}

	for ; l1 != nil; l1, l3 = l1.Next, l3.Next {
		sum := l1.Val + carry
		l3.Next = &ListNode{
			Val:  sum % 10,
		}
		carry = sum / 10
	}

	if carry != 0 {
		l3.Next = &ListNode{
			Val:  carry,
		}
	}

	return l3Head.Next
}

/*
Example:

Input: (3 -> 4 -> 2) + (4 -> 6 -> 5)
Output: 8 -> 0 -> 7
Explanation: 342 + 465 = 807.
 */

func addTwoNumbersReverse(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	l3, carry := calSum(l1, l2)
	if carry != 0 {
		return &ListNode{
			Val:  carry,
			Next: l3,
		}
	} else {
		return l3
	}
}

func calSum(l1 *ListNode, l2 *ListNode) (*ListNode, int) {
	if l1 == nil {
		return l2, 0
	}
	if l2 == nil {
		return l1, 0
	}

	l3, carry := calSum(l1.Next, l2.Next)
	sum := l1.Val + l2.Val + carry
	return &ListNode{
		Val:  sum % 10,
		Next: l3,
	}, sum / 10
}