package main

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,

return [0, 1].
 */

// twoSum
// time: O(n)
// storage: O(n)
func twoSum(nums []int, target int) []int {
	valueCache := make(map[int]int)

	for i, v := range nums {
		if adder, ok := valueCache[target-v]; ok {
			return []int{adder, i}
		} else {
			valueCache[v] = i
		}
	}

	return nil
}