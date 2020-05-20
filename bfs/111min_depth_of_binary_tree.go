package main

/*
Given a binary tree, find its minimum depth.

The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its minimum depth = 2.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	nowLength := 0

	for len(queue) > 0 {
		size := len(queue)
		nowLength++
		newQueue := make([]*TreeNode, 0)
		for i := 0; i < size; i++ {
			item := queue[i]

			if item.Left == nil && item.Right == nil {
				return nowLength
			}

			if item.Left != nil {
				newQueue = append(newQueue, item.Left)
			}
			if item.Right != nil {
				newQueue = append(newQueue, item.Right)
			}
		}
		queue = newQueue
	}
	return nowLength
}
