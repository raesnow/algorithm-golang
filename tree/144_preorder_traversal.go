package tree

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	result := []int{root.Val}
	left := preorderTraversal(root.Left)
	right := preorderTraversal(root.Right)
	result = append(result, left...)
	result = append(result, right...)
	return result
}
