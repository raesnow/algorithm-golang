package main

/*
Given an integer array nums, find the contiguous subarray (containing at least one number) which has the largest sum and return its sum.

Example:

Input: [-2,1,-3,4,-1,2,1,-5,4],
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
Follow up:

If you have figured out the O(n) solution, try coding another solution using the divide and conquer approach, which is more subtle.

*/

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	max := nums[0]
	cache := make([][]int, len(nums))
	for i := range cache {
		cache[i] = make([]int, len(nums))
		cache[i][i] = nums[i]
		if nums[i] > max {
			max = nums[i]
		}
	}

	for gap := 1; gap < len(nums); gap++ {
		for i := 0; i+gap < len(nums); i++ {
			cache[i][i+gap] = cache[i][i+gap-1] + nums[i+gap]
			if cache[i][i+gap] > max {
				max = cache[i][i+gap]
			}
		}
	}
	return max
}

func maxSubArray1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	copy(dp, nums)

	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] += dp[i-1]
		}
	}

	max := dp[0]
	for _, v := range dp {
		if v > max {
			max = v
		}
	}
	return max
}

func maxSubArray2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	pre, max := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		now := nums[i]
		if pre > 0 {
			now += pre
		}
		if now > max {
			max = now
		}
		pre = now
	}
	return max
}
