package main

/*
There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]
nums2 = [2]

The median is 2.0
Example 2:

nums1 = [1, 2]
nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5
*/

// findMedianSortedArrays
// time: O(log(min(m, n))
// storage: O(1)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) == 0 {
		return findMedian(nums2)
	}
	if len(nums2) == 0 {
		return findMedian(nums1)
	}

	length1 := len(nums1)
	length2 := len(nums2)
	if length1 > length2 {
		nums1, nums2 = nums2, nums1
		length1, length2 = length2, length1
	}

	half1 := length1 / 2
	half1Left := 0
	half1Right := length1
	for {
		half2 := (length1+length2)/2 - half1
		if (half1 == 0 && nums2[half2-1] <= nums1[half1]) || (half1 == length1 && nums1[half1-1] <= nums2[half2]) {
			return calMedian(nums1, nums2, half1, half2, length1, length2)
		}
		if (half2 == 0 && nums1[half1-1] <= nums2[half2]) || (half2 == length2 && nums2[half2-1] <= nums1[half1]) {
			return calMedian(nums1, nums2, half1, half2, length1, length2)
		}
		if half1 > 0 && nums1[half1-1] <= nums2[half2] && nums2[half2-1] <= nums1[half1] {
			return calMedian(nums1, nums2, half1, half2, length1, length2)
		}

		if half1 > 0 && nums1[half1-1] > nums2[half2] {
			half1Right = half1
		} else {
			half1Left = half1 + 1
		}
		half1 = (half1Left + half1Right) / 2
	}
}

func findMedian(nums []int) float64 {
	length := len(nums)
	if length == 0 {
		return 0
	}

	half := length / 2
	if length%2 == 0 {
		return float64(nums[half-1]+nums[half]) / 2.0
	} else {
		return float64(nums[half])
	}
}

func calMedian(nums1, nums2 []int, half1, half2, length1, length2 int) float64 {
	odd := true
	if (length1+length2)%2 == 0 {
		odd = false
	}

	if odd {
		leftCount := half1 + half2
		rightCount := length1 + length2 - half1 - half2
		if leftCount > rightCount {
			if half1 == 0 {
				return float64(nums2[half2-1])
			}
			if half2 == 0 {
				return float64(nums2[half1-1])
			}
			return float64(max(nums1[half1-1], nums2[half2-1]))
		} else {
			if half1 == length1 {
				return float64(nums2[half2])
			}
			if half2 == length2 {
				return float64(nums2[half1])
			}
			return float64(min(nums1[half1], nums2[half2]))
		}
	} else {
		left := 0
		if half1 == 0 {
			left = nums2[half2-1]
		} else if half2 == 0 {
			left = nums1[half1-1]
		} else {
			left = max(nums1[half1-1], nums2[half2-1])
		}

		right := 0
		if half1 == length1 {
			right = nums2[half2]
		} else if half2 == length2 {
			right = nums1[half1]
		} else {
			right = min(nums1[half1], nums2[half2])
		}
		return float64(left+right) / 2.0
	}
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

func min(a, b int) int {
	if a <= b {
		return a
	} else {
		return b
	}
}
