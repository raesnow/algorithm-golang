package array

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}

	left, right := 0, 0
	result := 0
	windowMap := make(map[uint8]int)
	for ; right < len(s); right++ {
		windowMap[s[right]]++

		for windowMap[s[right]] > 1 {
			windowMap[s[left]]--
			left++
		}
		windowLen := right - left + 1
		if windowLen > result {
			result = windowLen
		}
	}
	return result
}
