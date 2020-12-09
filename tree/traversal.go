package main

func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	left := preorderTraversal(root.Left)
	right := preorderTraversal(root.Right)

	result := []int{root.Val}
	result = append(result, left...)
	result = append(result, right...)
	return result
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	left := inorderTraversal(root.Left)
	right := inorderTraversal(root.Right)

	result := left
	result = append(result, root.Val)
	result = append(result, right...)
	return result
}

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	left := postorderTraversal(root.Left)
	right := postorderTraversal(root.Right)

	result := left
	result = append(result, right...)
	result = append(result, root.Val)
	return result
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	queue := []*TreeNode{root}
	result := make([][]int, 0)
	for len(queue) != 0 {
		newQueue := make([]*TreeNode, 0)
		curLevel := make([]int, 0)
		for _, v := range queue {
			if v != nil {
				curLevel = append(curLevel, v.Val)
				newQueue = append(newQueue, v.Left)
				newQueue = append(newQueue, v.Right)
			}
		}
		if len(curLevel) != 0 {
			result = append(result, curLevel)
		}
		queue = newQueue
	}
	return result
}
