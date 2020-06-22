package main

import "fmt"

/*
Given a collection of distinct integers, return all possible permutations.

Example:

Input: [1,2,3]
Output:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

*/

var result [][]int

func permute(nums []int) [][]int {
	result = make([][]int, 0)

	if len(nums) == 0 {
		return result
	}
	path := make([]int, 0)
	backtrack(path, nums, len(nums))
	return result
}

func backtrack(path, nums []int, endLength int) {
	if len(path) == endLength {
		fmt.Printf("add path: %#v\n", path)
		now := make([]int, len(path))
		copy(now, path)
		result = append(result, now)
		return
	}
	for i, v := range nums {
		var choices []int
		if i == 0 {
			choices = nums[1:]
		} else {
			choices = append(choices, nums[0:i]...)
			choices = append(choices, nums[i+1:]...)
		}
		path = append(path, v)
		fmt.Printf("path: %#v, choices: %#v\n", path, choices)
		backtrack(path, choices, endLength)
		path = path[:len(path)-1]
	}
}
