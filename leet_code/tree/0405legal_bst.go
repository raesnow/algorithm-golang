package main

/*
Implement a function to check if a binary tree is a binary search tree.

Example 1:

Input:
    2
   / \
  1   3
Output: true
Example 2:

Input:
    5
   / \
  1   4
     / \
    3   6
Output: false
Explanation: Input: [5,1,4,null,null,3,6].
     the value of root node is 5, but its right child has value 4.
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	status, _, _ := checkValidBST(root)
	return status
}

func checkValidBST(root *TreeNode) (isvalid bool, min, max int) {
	if root.Right == nil && root.Left == nil {
		return true, root.Val, root.Val
	}
	if root.Right == nil {
		if isValid, min, max := checkValidBST(root.Left); isValid {
			if root.Val > max {
				return true, min, root.Val
			}
		}
		return false, 0, 0
	}
	if root.Left == nil {
		if isValid, min, max := checkValidBST(root.Right); isValid {
			if root.Val < min {
				return true, root.Val, max
			}
		}
		return false, 0, 0
	}

	leftIsValid, leftMin, leftMax := checkValidBST(root.Left)
	rightIsValid, rightMin, rightMax := checkValidBST(root.Right)
	if leftIsValid && rightIsValid && root.Val > leftMax && root.Val < rightMin {
		return true, leftMin, rightMax
	}
	return false, 0, 0
}
