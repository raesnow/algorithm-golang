package main

/*
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.

Example 2:

Input: "cbbd"
Output: "bb"
*/

const (
	UNKNOWN = iota
	TRUE
	FALSE
)

// longestPalindrome
// time: O(n*n)
// storage: O(n*n)
func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	isPalindromeCache := make([][]int, len(s))
	for i := range isPalindromeCache {
		isPalindromeCache[i] = make([]int, len(s))
		isPalindromeCache[i][i] = TRUE
	}

	longestLength := 1
	longestSubStr := s[0:1]
	for gap := 1; gap < len(s); gap++ {
		for i := 0; i+gap < len(s); i++ {
			j := i + gap
			if gap == 1 {
				if s[i] == s[j] {
					isPalindromeCache[i][j] = TRUE
				} else {
					isPalindromeCache[i][j] = FALSE
					continue
				}
			} else {
				if isPalindromeCache[i+1][j-1] == TRUE && s[i] == s[j] {
					isPalindromeCache[i][j] = TRUE
				} else {
					isPalindromeCache[i][j] = FALSE
					continue
				}
			}

			if gap+1 > longestLength {
				longestLength = gap + 1
				longestSubStr = s[i : j+1]
			}
		}
	}

	return longestSubStr
}

// longestPalindrome1
// time: O(n*n)
// storage: O(1)
func longestPalindrome1(s string) string {
	if len(s) == 0 {
		return ""
	}

	longestLength := 1
	longestStr := s[0:1]
	for i := range s {
		length1, result1 := scale(s, i, -1)
		length2, result2 := scale(s, i, i+1)
		if length1 >= length2 && length1 > longestLength {
			longestLength = length1
			longestStr = result1
		} else if length1 < length2 && length2 > longestLength {
			longestLength = length2
			longestStr = result2
		}
	}
	return longestStr
}

func scale(s string, i, j int) (length int, result string) {
	var left, right int
	if j == -1 {
		left, right = i-1, i+1
	} else {
		left, right = i, j
	}
	for left >= 0 && right < len(s) {
		if s[left] == s[right] {
			length = right - left + 1
			result = s[left : right+1]
		} else {
			return
		}
		left--
		right++
	}
	return
}

// longestPalindrome2
// time: O(n)
// storage: O(n)
func longestPalindrome2(s string) string {
	if len(s) == 0 || len(s) == 1 {
		return s
	}

	manacherStr := manacherString(s)
	radius := make([]int, len(manacherStr))
	R, c, max, index := -1, -1, 0, 0
	for i := 0; i < len(manacherStr); i++ {
		if R > i {
			if radius[2*c-i] <= R-i+1 {
				radius[i] = radius[2*c-i]
			} else {
				radius[i] = R - i + 1
			}
		} else {
			radius[i] = 1
		}

		for i+radius[i] < len(manacherStr) && i-radius[i] >= 0 {
			if manacherStr[i-radius[i]] == manacherStr[i+radius[i]] {
				radius[i]++
			} else {
				break
			}
		}

		if i+radius[i]-1 > R {
			R = i + radius[i] - 1
			c = i
		}

		if radius[i] > max {
			max = radius[i]
			index = i
		}
	}

	result := make([]byte, 0)
	for i := index - max + 1; i < index+max; i++ {
		if manacherStr[i] != '#' {
			result = append(result, manacherStr[i])
		}
	}
	return string(result)
}

func manacherString(s string) string {
	result := make([]byte, 2*len(s)+1)
	j := 0
	for _, v := range []byte(s) {
		result[j] = byte('#')
		j++
		result[j] = v
		j++
	}
	result[j] = byte('#')
	return string(result)
}
