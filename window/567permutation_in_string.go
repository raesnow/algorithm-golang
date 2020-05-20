package main

/*
Given two strings s1 and s2, write a function to return true if s2 contains the permutation of s1. In other words, one of the first string's permutations is the substring of the second string.



Example 1:

Input: s1 = "ab" s2 = "eidbaooo"
Output: True
Explanation: s2 contains one permutation of s1 ("ba").
Example 2:

Input:s1= "ab" s2 = "eidboaoo"
Output: False


Note:

The input strings only contain lower case letters.
The length of both given strings is in range [1, 10,000].
*/

func checkInclusion(s1 string, s2 string) bool {
	needs := make(map[uint8]int)
	includes := make(map[uint8]int)
	for i := range s1 {
		if _, ok := needs[s1[i]]; !ok {
			needs[s1[i]] = 1
		} else {
			needs[s1[i]]++
		}
		includes[s1[i]] = 0
	}

	left, right := 0, 0
	valid := 0
	for right < len(s2) {
		v := s2[right]
		if _, ok := needs[v]; ok {
			includes[v]++
			if includes[v] == needs[v] {
				valid++
			}
		}
		right++

		for right-left >= len(s1) && valid == len(needs) {
			if right-left == len(s1) {
				return true
			}

			v := s2[left]
			if _, ok := needs[v]; ok {
				if includes[v] == needs[v] {
					valid--
				}
				includes[v]--
			}
			left++
		}
	}
	return false
}
