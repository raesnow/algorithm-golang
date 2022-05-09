package main

var lis []int

func lengthOfLIS2(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	lis = make([]int, len(nums))
	lengthOfLISDp(nums)

	max := 0
	for _, v := range lis {
		if v > max {
			max = v
		}
	}
	return max
}

func lengthOfLISDp(nums []int) {
	for i := range nums {
		preMax := 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && lis[j]+1 > preMax {
				preMax = lis[j] + 1
			}
		}
		lis[i] = preMax
	}
}
