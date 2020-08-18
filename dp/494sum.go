package main

import "fmt"

/*
You are given a list of non-negative integers, a1, a2, ..., an, and a target, S. Now you have 2 symbols + and -. For each integer, you should choose one from + and - as its new symbol.

Find out how many ways to assign symbols to make sum of integers equal to target S.

Example 1:

Input: nums is [1, 1, 1, 1, 1], S is 3.
Output: 5
Explanation:

-1+1+1+1+1 = 3
+1-1+1+1+1 = 3
+1+1-1+1+1 = 3
+1+1+1-1+1 = 3
+1+1+1+1-1 = 3

There are 5 ways to assign symbols to make the sum of nums be target 3.


Constraints:

The length of the given array is positive and will not exceed 20.
The sum of elements in the given array will not exceed 1000.
Your output answer is guaranteed to be fitted in a 32-bit integer.

*/

var sumWaysCache = make(map[string]int)

func findTargetSumWays(nums []int, S int) int {
	sumWaysCache = make(map[string]int)
	if len(nums) == 0 {
		return 0
	}
	return backtrack(0, nums, S)
}

func backtrack(i int, nums []int, rest int) int {
	if i == len(nums) {
		if rest == 0 {
			return 1
		} else {
			return 0
		}
	}

	key := fmt.Sprintf("%d-%d", i, rest)
	if v, ok := sumWaysCache[key]; ok {
		return v
	}

	result := backtrack(i+1, nums, rest-nums[i]) + backtrack(i+1, nums, rest+nums[i])
	sumWaysCache[key] = result

	return result
}

func findTargetSumWays1(nums []int, S int) int {
	if len(nums) == 0 {
		return 0
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}
	if (sum < S) || ((sum+S)%2 == 1) {
		return 0
	}
	return subsets(nums, (sum+S)/2)
}

func subsets(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= len(nums); i++ {
		for j := target; j >= 0; j-- {
			if j >= nums[i-1] {
				dp[j] = dp[j] + dp[j-nums[i-1]]
			}
		}
	}
	return dp[target]
}
