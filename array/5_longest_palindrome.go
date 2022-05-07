package array

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}

	maxStr := ""
	for i := 0; i < len(s)-1; i++ {
		str1 := checkPalindrome(s, i, i)
		str2 := checkPalindrome(s, i, i+1)
		if len(str1) > len(maxStr) {
			maxStr = str1
		}
		if len(str2) > len(maxStr) {
			maxStr = str2
		}
	}
	return maxStr
}

func checkPalindrome(s string, left, right int) string {
	for left >= 0 && right < len(s) {
		if s[left] != s[right] {
			break
		}
		left--
		right++
	}
	if left == right {
		return s[left:right]
	}
	return s[left+1 : right]
}
