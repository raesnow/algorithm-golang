package main

/*
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:

[
  [3],
  [9,20],
  [15,7]
]
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	currentLevel := 1
	nextLevel := 0

	result := make([][]int, 0)
	for len(queue) > 0 {
		level := make([]int, 0)
		for ; currentLevel > 0; currentLevel-- {
			if queue[0] == nil {
				queue = queue[1:]
				continue
			}
			level = append(level, queue[0].Val)
			queue = append(queue, queue[0].Left)
			queue = append(queue, queue[0].Right)
			queue = queue[1:]
			nextLevel += 2
		}
		if len(level) == 0 {
			break
		}
		result = append(result, level)
		currentLevel = nextLevel
		nextLevel = 0
	}
	return result
}
