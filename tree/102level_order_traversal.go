package main

//Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).
//
//For example:
//Given binary tree [3,9,20,null,null,15,7],
//    3
//   / \
//  9  20
//    /  \
//   15   7
//return its level order traversal as:
//[
//  [3],
//  [9,20],
//  [15,7]
//]
//

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
