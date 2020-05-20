package main

/*
Given a string s and a non-empty string p, find all the start indices of p's anagrams in s.

Strings consists of lowercase English letters only and the length of both strings s and p will not be larger than 20,100.

The order of output does not matter.

Example 1:

Input:
s: "cbaebabacd" p: "abc"

Output:
[0, 6]

Explanation:
The substring with start index = 0 is "cba", which is an anagram of "abc".
The substring with start index = 6 is "bac", which is an anagram of "abc".
Example 2:

Input:
s: "abab" p: "ab"

Output:
[0, 1, 2]

Explanation:
The substring with start index = 0 is "ab", which is an anagram of "ab".
The substring with start index = 1 is "ba", which is an anagram of "ab".
The substring with start index = 2 is "ab", which is an anagram of "ab".
*/

func findAnagrams(s string, p string) []int {
	needs := make(map[uint8]int)
	includes := make(map[uint8]int)
	for i := range p {
		if _, ok := needs[p[i]]; !ok {
			needs[p[i]] = 1
		} else {
			needs[p[i]]++
		}
		includes[p[i]] = 0
	}

	left, right := 0, 0
	valid := 0
	result := make([]int, 0)
	for right < len(s) {
		v := s[right]
		if _, ok := includes[v]; ok {
			includes[v]++
			if includes[v] == needs[v] {
				valid++
			}
		}
		right++

		for right-left >= len(p) && valid == len(needs) {
			if right-left == len(p) {
				result = append(result, left)
				break
			}

			v := s[left]
			if _, ok := includes[v]; ok {
				if includes[v] == needs[v] {
					valid--
				}
				includes[v]--
			}
			left++
		}
	}
	return result
}
