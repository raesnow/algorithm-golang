package array

func removeDuplicates(nums []int) int {
	if len(nums) == 0 || len(nums) == 1 {
		return len(nums)
	}

	slow, fast := 0, 1
	for ; fast < len(nums); fast++ {
		if nums[fast] != nums[slow] {
			slow += 1
			nums[slow] = nums[fast]
		}
	}
	return slow + 1
}
