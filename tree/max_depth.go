package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := maxDepth(root.Left)
	right := maxDepth(root.Right)

	result := left
	if result < right {
		result = right
	}
	return result + 1
}
