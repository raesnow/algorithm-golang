package tree

func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	traverse(root, k)
	return result
}

var rank = 0
var result = 0

func traverse(root *TreeNode, k int) {
	if root == nil {
		return
	}

	traverse(root.Left, k)
	rank++
	if rank == k {
		result = root.Val
		return
	}
	traverse(root.Right, k)
}
