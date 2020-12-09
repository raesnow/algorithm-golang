package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return equal(root.Left, root.Right)
}

func equal(node1, node2 *TreeNode) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	if node1.Val != node2.Val {
		return false
	}

	if equal(node1.Left, node2.Right) && equal(node1.Right, node2.Left) {
		return true
	}
	return false
}
