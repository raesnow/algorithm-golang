package main

/*
Given a non-empty array containing only positive integers,
find if the array can be partitioned into two subsets such that the sum of elements in both subsets
is equal.

Note:

Each of the array element will not exceed 100.
The array size will not exceed 200.


Example 1:

Input: [1, 5, 11, 5]

Output: true

Explanation: The array can be partitioned as [1, 5, 5] and [11].


Example 2:

Input: [1, 2, 3, 5]

Output: false

Explanation: The array cannot be partitioned into equal sum subsets.
*/

func canPartition1(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return false
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}

	cache := make([][]bool, len(nums)+1)
	for i := range cache {
		cache[i] = make([]bool, sum/2+1)
		cache[i][0] = true
	}

	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= sum/2; j++ {
			cache[i][j] = cache[i-1][j]
			if j >= nums[i-1] {
				cache[i][j] = cache[i][j] || cache[i-1][j-nums[i-1]]
			}
		}
	}
	return cache[len(nums)][sum/2]
}

func canPartition(nums []int) bool {
	if len(nums) == 0 || len(nums) == 1 {
		return false
	}

	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}

	cache := make([]bool, sum/2+1)
	cache[0] = true

	for i := 1; i <= len(nums); i++ {
		for j := sum / 2; j > 0; j-- {
			if j >= nums[i-1] {
				cache[j] = cache[j] || cache[j-nums[i-1]]
			}
		}
	}
	return cache[sum/2]
}
