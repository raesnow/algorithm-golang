package array

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}
	tMap := make(map[uint8]int)
	for i := range t {
		tMap[t[i]]++
	}

	left, right := 0, 0
	windowMap := make(map[uint8]int)
	result := ""
	for ; right < len(s); right++ {
		windowMap[s[right]]++

		for checkWindow(tMap, windowMap) {
			if len(result) == 0 || right-left+1 < len(result) {
				result = s[left : right+1]
			}
			windowMap[s[left]]--
			left++
		}
	}
	return result
}

func checkWindow(origin, window map[uint8]int) bool {
	if len(origin) > len(window) {
		return false
	}
	for k, v := range origin {
		vWindow, ok := window[k]
		if !ok {
			return false
		}
		if vWindow < v {
			return false
		}
	}
	return true
}
