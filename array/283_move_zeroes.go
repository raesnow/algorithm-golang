package array

func moveZeroes(nums []int) []int {
	if len(nums) == 0 {
		return nums
	}

	slow, fast := 0, 0
	for ; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow] = nums[fast]
			slow++
		}
	}
	for ; slow < len(nums); slow++ {
		nums[slow] = 0
	}
	return nums
}
