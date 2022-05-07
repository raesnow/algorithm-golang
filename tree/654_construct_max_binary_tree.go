package tree

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	maxIndex := 0
	maxVal := nums[0]
	for i := range nums {
		if nums[i] > maxVal {
			maxIndex = i
			maxVal = nums[i]
		}
	}

	var left *TreeNode
	if maxIndex > 0 {
		left = constructMaximumBinaryTree(nums[:maxIndex])
	}
	var right *TreeNode
	if maxIndex < len(nums)-1 {
		right = constructMaximumBinaryTree(nums[maxIndex+1:])
	}
	return &TreeNode{
		Val:   maxVal,
		Left:  left,
		Right: right,
	}
}
