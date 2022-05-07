package tree

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}

	rootVal := preorder[0]
	inorderRootIndex := 0
	for i := range inorder {
		if inorder[i] == rootVal {
			inorderRootIndex = i
			break
		}
	}
	var left *TreeNode
	if inorderRootIndex > 0 {
		left = buildTree(preorder[1:inorderRootIndex+1], inorder[0:inorderRootIndex])
	}
	var right *TreeNode
	if inorderRootIndex < len(inorder)-1 {
		right = buildTree(preorder[inorderRootIndex+1:], inorder[inorderRootIndex+1:])
	}
	return &TreeNode{
		Val:   rootVal,
		Left:  left,
		Right: right,
	}
}
