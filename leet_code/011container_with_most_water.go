package main

/*
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.

Example:

Input: [1,8,6,2,5,4,8,3,7]
Output: 49
*/

// maxArea
// time: O(n)
// storage: O(1)
func maxArea(height []int) int {
	maxValue := 0
	for i, j := 0, len(height)-1; i < j; {
		iIsMin := true
		minHeight := height[i]
		if height[i] > height[j] {
			minHeight = height[j]
			iIsMin = false
		}

		value := minHeight * (j - i)
		if value > maxValue {
			maxValue = value
		}

		if iIsMin {
			i++
		} else {
			j--
		}
	}
	return maxValue
}
