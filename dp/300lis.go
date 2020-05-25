package main

/*
Given an unsorted array of integers, find the length of longest increasing subsequence.

Example:

Input: [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
Note:

There may be more than one LIS combination, it is only necessary for you to return the length.
Your algorithm should run in O(n2) complexity.
Follow up: Could you improve it to O(n log n) time complexity?
*/

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	for i := range nums {
		max := 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && dp[j]+1 > max {
				max = dp[j] + 1
			}
		}
		dp[i] = max
	}

	result := 1
	for _, v := range dp {
		if v > result {
			result = v
		}
	}
	return result
}

func lengthOfLIS1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	top := make([]int, 0)

	for i := range nums {
		value := nums[i]

		left, right := 0, len(top)
		for left < right {
			mid := left + (right-left)/2
			if top[mid] == value {
				right = mid
			} else if top[mid] > value {
				right = mid
			} else {
				left = mid + 1
			}
		}

		if left == len(top) {
			top = append(top, value)
		}
		top[left] = value
	}
	return len(top)
}
