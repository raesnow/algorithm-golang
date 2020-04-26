package main

/*
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]
*/

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	result := make([][]int, 0)

	for i, v := range nums {
		twoSumResult := calTwoSum(nums, 0-v, i)
		for _, sum := range twoSumResult {
			threes := sort([]int{v, sum[0], sum[1]})
			result = deDuplicate(result, threes)
		}
	}

	return result
}

func calTwoSum(nums []int, sum int, skipIndex int) [][]int {
	cache := make(map[int]int)
	result := make([][]int, 0)

	for i, v := range nums {
		if i == skipIndex {
			continue
		}
		if _, ok := cache[sum-v]; ok {
			result = append(result, []int{sum - v, v})
		}
		cache[v] = i
	}
	return result
}

func sort(threes []int) []int {
	for i := 0; i < len(threes)-1; i++ {
		for j := 0; j < len(threes)-i-1; j++ {
			if threes[j] > threes[j+1] {
				threes[j], threes[j+1] = threes[j+1], threes[j]
			}
		}
	}
	return threes
}

func deDuplicate(result [][]int, threes []int) [][]int {
	if len(result) == 0 {
		result = append(result, threes)
		return result
	}

	for _, v := range result {
		if threes[0] == v[0] && threes[1] == v[1] && threes[2] == v[2] {
			return result
		}
	}
	result = append(result, threes)
	return result
}
