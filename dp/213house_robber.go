package main

/*
You are a professional robber planning to rob houses along a street.
Each house has a certain amount of money stashed. All houses at this place are arranged in a circle.
That means the first house is the neighbor of the last one.
Meanwhile, adjacent houses have security system connected and it will automatically contact the police
if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house,
determine the maximum amount of money you can rob tonight without alerting the police.

Example 1:

Input: [2,3,2]
Output: 3
Explanation: You cannot rob house 1 (money = 2) and then rob house 3 (money = 2),
             because they are adjacent houses.
Example 2:

Input: [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
             Total amount you can rob = 1 + 3 = 4.
*/

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	robFirst := rob2Dp(nums, 0, len(nums)-2)
	robLast := rob2Dp(nums, 1, len(nums)-1)
	result := robFirst
	if robLast > robFirst {
		result = robLast
	}
	return result
}

func rob2Dp(nums []int, start, end int) int {
	if start >= len(nums) {
		return 0
	}
	if end < 0 {
		return 0
	}

	pre, prePre := 0, 0
	result := 0
	for i := start; i <= end; i++ {
		do := nums[i] + prePre
		notDo := pre
		result = do
		if notDo > do {
			result = notDo
		}

		prePre = pre
		pre = result
	}
	return result
}
