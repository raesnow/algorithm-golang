package array

type NumArray struct {
	nums    []int
	preSums []int
}

func Constructor(nums []int) NumArray {
	preSums := make([]int, len(nums)+1)

	preSums[0] = 0
	for i, value := range nums {
		preSums[i+1] = preSums[i] + value
	}
	return NumArray{
		nums:    nums,
		preSums: preSums,
	}
}

func (this *NumArray) SumRange(left int, right int) int {
	if left > right || left < 0 || len(this.nums) == 0 {
		return 0
	}
	if right >= len(this.preSums) {
		right = len(this.preSums) - 1
	}
	return this.preSums[right+1] - this.preSums[left]
}
