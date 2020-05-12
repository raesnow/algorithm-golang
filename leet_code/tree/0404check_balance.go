package main

/*
Implement a function to check if a binary tree is balanced. For the purposes of this question,
a balanced tree is defined to be a tree such that the heights of the two subtrees of any node never differ by more than one.


Example 1:

Given tree [3,9,20,null,null,15,7]
    3
   / \
  9  20
    /  \
   15   7
return true.
Example 2:

Given [1,2,2,3,3,null,null,4,4]
      1
     / \
    2   2
   / \
  3   3
 / \
4   4
return false.
*/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if isBalanced(root.Left) && isBalanced(root.Right) {
		diff := getHeight(root.Left) - getHeight(root.Right)
		if diff == 0 || diff == -1 || diff == 1 {
			return true
		} else {
			return false
		}
	}
	return false
}

var heightMap = make(map[*TreeNode]int)

func getHeight(root *TreeNode) int {
	if val, ok := heightMap[root]; ok {
		return val
	}
	if root == nil {
		heightMap[root] = 0
		return 0
	}

	leftHeight := getHeight(root.Left)
	rightHeight := getHeight(root.Right)
	max := leftHeight
	if rightHeight > leftHeight {
		max = rightHeight
	}
	heightMap[root] = max + 1
	return max + 1
}
