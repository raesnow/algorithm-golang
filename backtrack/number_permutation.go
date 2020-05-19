package main

/*
 穷举全排列,比方说给三个数 [1,2,3]
*/

var result = make([][]int, 0)

func permutation(nums []int) [][]int {
	if len(nums) == 0 {
		return result
	}
	path := make([]int, 0)
	backtrack(path, nums, len(nums))
	return result
}

func backtrack(path, nums []int, endLength int) {
	if len(path) == endLength {
		result = append(result, path)
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
		backtrack(path, choices, endLength)
		path = path[:len(path)-1]
	}
}
