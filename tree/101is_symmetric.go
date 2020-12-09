package main

//Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).
//
//For example, this binary tree [1,2,2,3,4,4,3] is symmetric:
//
//    1
//   / \
//  2   2
// / \ / \
//3  4 4  3
//
//But the following [1,2,2,null,3,null,3] is not:
//
//    1
//   / \
//  2   2
//   \   \
//   3    3
//
//Follow up: Solve it both recursively and iteratively.
//

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
