package main

func buildTree1(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 && len(postorder) == 0 {
		return nil
	}
	if len(inorder) == 1 {
		return &TreeNode{Val: inorder[0]}
	}

	rootVal := postorder[len(postorder)-1]
	root := &TreeNode{
		Val: rootVal,
	}

	rootIndex := 0
	for i, v := range inorder {
		if v == rootVal {
			rootIndex = i
		}
	}
	root.Left = buildTree(inorder[0:rootIndex], postorder[0:rootIndex])
	root.Right = buildTree(inorder[rootIndex+1:], postorder[rootIndex:len(postorder)-1])
	return root
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 && len(inorder) == 0 {
		return nil
	}
	if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}

	rootVal := preorder[0]
	root := &TreeNode{
		Val: rootVal,
	}

	rootIndex := 0
	for i, v := range inorder {
		if v == rootVal {
			rootIndex = i
		}
	}
	root.Left = buildTree(preorder[1:rootIndex+1], inorder[0:rootIndex])
	root.Right = buildTree(preorder[rootIndex+1:], inorder[rootIndex+1:])
	return root
}
