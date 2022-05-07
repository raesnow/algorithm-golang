package sort

func sortArray4(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	chooseSort(nums)
	return nums
}

func chooseSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	for i := 0; i < len(nums); i++ {
		minVal := nums[i]
		minIndex := i
		for j := i; j < len(nums); j++ {
			if nums[j] < minVal {
				minIndex = j
				minVal = nums[j]
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}
