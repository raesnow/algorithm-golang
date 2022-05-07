package sort

import "math/rand"

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, start, end int) {
	if start >= end {
		return
	}
	pivot := quickSortPartition(nums, start, end)
	quickSort(nums, start, pivot)
	quickSort(nums, pivot+1, end)
	return
}

func quickSortPartition(nums []int, start, end int) int {
	pivotIndex := rand.Intn(end-start+1) + start
	pivotVal := nums[pivotIndex]
	nums[pivotIndex], nums[end] = nums[end], nums[pivotIndex]

	left, right := start-1, start
	for ; right < end; right++ {
		if nums[right] < pivotVal {
			left++
			nums[left], nums[right] = nums[right], nums[left]
		}
	}
	nums[left+1], nums[end] = nums[end], nums[left+1]
	return left + 1
}
