package main

/*
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.
*/

// lengthOfLongestSubstring
// time: O(n)
// storage: O(min(m, n)), 字符集为m
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}

	charLocation := make(map[byte]int)

	head := 0
	charLocation[s[head]] = head
	maxLength := 1
	for iNew, v := range []byte(s[1:]) {
		i := iNew + 1
		if loc, ok := charLocation[v]; ok {
			if loc >= head {
				head = loc + 1
				charLocation[v] = i
				continue
			}
		}
		charLocation[v] = i
		length := i - head + 1
		if length > maxLength {
			maxLength = length
		}
	}

	return maxLength
}
