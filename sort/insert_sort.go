package sort

func sortArray3(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	insertSort(nums)
	return nums
}

func insertSort(nums []int) {
	if len(nums) <= 1 {
		return
	}

	for i := 1; i < len(nums); i++ {
		value := nums[i]
		for j := i - 1; j >= 0; j-- {
			if nums[j] > value {
				nums[j+1] = nums[j]
			} else {
				nums[j+1] = value
				break
			}
		}
	}
}
