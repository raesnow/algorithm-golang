package main

/*
The thief has found himself a new place for his thievery again. There is only one entrance to this area, called the "root." Besides the root, each house has one and only one parent house. After a tour, the smart thief realized that "all houses in this place forms a binary tree". It will automatically contact the police if two directly-linked houses were broken into on the same night.

Determine the maximum amount of money the thief can rob tonight without alerting the police.

Example 1:

Input: [3,2,3,null,3,null,1]

     3
    / \
   2   3
    \   \
     3   1

Output: 7
Explanation: Maximum amount of money the thief can rob = 3 + 3 + 1 = 7.
Example 2:

Input: [3,4,5,1,3,null,1]

     3
    / \
   4   5
  / \   \
 1   3   1

Output: 9
Explanation: Maximum amount of money the thief can rob = 4 + 5 = 9.
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var rob3Cache map[*TreeNode]int

func rob3(root *TreeNode) int {
	rob3Cache = make(map[*TreeNode]int)
	return rob3Dp(root)
}

func rob3Dp(root *TreeNode) int {
	if root == nil {
		return 0
	}

	if v, ok := rob3Cache[root]; ok {
		return v
	}

	doRightGrandsons, doLeftGrandsons := 0, 0
	if root.Right != nil {
		doRightGrandsons = rob3Dp(root.Right.Right) + rob3Dp(root.Right.Left)
	}
	if root.Left != nil {
		doLeftGrandsons = rob3Dp(root.Left.Right) + rob3Dp(root.Left.Left)
	}

	do := root.Val + doRightGrandsons + doLeftGrandsons
	notDo := rob3Dp(root.Right) + rob3Dp(root.Left)
	result := do
	if notDo > do {
		result = notDo
	}
	rob3Cache[root] = result
	return result
}

func rob4(root *TreeNode) int {
	do, notDo := rob4Dp(root)
	result := do
	if notDo > do {
		result = notDo
	}
	return result
}

func rob4Dp(root *TreeNode) (int, int) {
	if root == nil {
		return 0, 0
	}

	leftDo, leftNotDo := rob4Dp(root.Left)
	rightDo, rightNotDo := rob4Dp(root.Right)
	do := root.Val + leftNotDo + rightNotDo
	leftMax := leftDo
	if leftNotDo > leftMax {
		leftMax = leftNotDo
	}
	rightMax := rightDo
	if rightNotDo > rightMax {
		rightMax = rightNotDo
	}
	notDo := leftMax + rightMax
	return do, notDo
}
