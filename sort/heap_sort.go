package sort

func sortArray1(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	heapSort(nums)
	return nums
}

func heapSort(nums []int) {
	buildHeap(nums)

	length := len(nums)
	for i := len(nums) - 1; i >= 0; i-- {
		nums[i], nums[0] = nums[0], nums[i]
		length--
		maxHeapify(nums, 0, length)
	}
}

func buildHeap(nums []int) {
	for i := len(nums) / 2; i >= 0; i-- {
		maxHeapify(nums, i, len(nums))
	}
}

func maxHeapify(nums []int, i int, length int) {
	for i*2+1 < length {
		lson := i*2 + 1
		rson := i*2 + 2
		largeIndex := i
		if lson < length && nums[lson] > nums[i] {
			largeIndex = lson
		}
		if rson < length && nums[rson] > nums[largeIndex] {
			largeIndex = rson
		}
		if i == largeIndex {
			break
		}
		nums[i], nums[largeIndex] = nums[largeIndex], nums[i]
		i = largeIndex
	}
}
