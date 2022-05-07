package sort

func sortArray2(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	heapSort(nums)
	return nums
}

func mergeSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := mergeSort(nums[:mid])
	right := mergeSort(nums[mid:])

	return merge(left, right)
}

func merge(array1, array2 []int) []int {
	if len(array1) == 0 {
		return array2
	}
	if len(array2) == 0 {
		return array1
	}

	index1, index2 := 0, 0
	index := 0
	result := make([]int, len(array1)+len(array2))
	for index1 < len(array1) && index2 < len(array2) {
		if array1[index1] < array2[index2] {
			result[index] = array1[index1]
			index1++
		} else {
			result[index] = array2[index2]
			index2++
		}
		index++
	}

	for index1 < len(array1) {
		result[index] = array1[index1]
		index1++
		index++
	}
	for index2 < len(array2) {
		result[index] = array2[index2]
		index2++
		index++
	}
	return result
}
