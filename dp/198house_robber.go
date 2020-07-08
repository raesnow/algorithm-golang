package main

/*
You are a professional robber planning to rob houses along a street.
Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that
adjacent houses have security system connected and it will automatically contact the police
if two adjacent houses were broken into on the same night.

Given a list of non-negative integers representing the amount of money of each house,
determine the maximum amount of money you can rob tonight without alerting the police.



Example 1:

Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
             Total amount you can rob = 1 + 3 = 4.
Example 2:

Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
             Total amount you can rob = 2 + 9 + 1 = 12.


Constraints:

0 <= nums.length <= 100
0 <= nums[i] <= 400

*/

var res []int

func rob(nums []int) int {
	res = make([]int, len(nums))
	return robDp(nums, 0)
}

func robDp(nums []int, start int) int {
	if start >= len(nums) {
		return 0
	}

	if res[start] != 0 {
		return res[start]
	}

	do := nums[start] + robDp(nums, start+2)
	notDo := robDp(nums, start+1)

	if do >= notDo {
		res[start] = do
	} else {
		res[start] = notDo
	}
	return res[start]
}

func rob1(nums []int) int {
	pre, prePre := 0, 0
	result := 0
	for _, v := range nums {
		do := v + prePre
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
