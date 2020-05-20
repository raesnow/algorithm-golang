package main

/*
Given a string S and a string T, find the minimum window in S which will contain all the characters in T in complexity O(n).

Example:

Input: S = "ADOBECODEBANC", T = "ABC"
Output: "BANC"
Note:

If there is no such window in S that covers all characters in T, return the empty string "".
If there is such window, you are guaranteed that there will always be only one unique minimum window in S.
*/

func minWindow(s string, t string) string {
	if len(s) == 0 || len(t) == 0 {
		return ""
	}

	needs := make(map[uint8]int)
	includes := make(map[uint8]int)
	for i := range t {
		if _, ok := needs[t[i]]; !ok {
			needs[t[i]] = 1
		} else {
			needs[t[i]]++
		}
		includes[t[i]] = 0
	}

	left, right := 0, 0
	valid := 0
	minLength, start := 0, 0
	for right < len(s) {
		if _, ok := includes[s[right]]; ok {
			includes[s[right]]++
			if includes[s[right]] == needs[s[right]] {
				valid++
			}
		}
		right++

		for valid == len(needs) {
			if minLength == 0 || minLength > right-left {
				minLength = right - left
				start = left
			}
			if _, ok := includes[s[left]]; ok {
				if includes[s[left]] == needs[s[left]] {
					valid--
				}
				includes[s[left]]--
			}
			left++
		}
	}
	if minLength == 0 {
		return ""
	}
	return s[start : start+minLength]
}
