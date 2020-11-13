package main

/*
Given a string s, find the longest palindromic subsequence's length in s. You may assume that the maximum length of s is 1000.

Example 1:
Input:

"bbbab"
Output:
4
One possible longest palindromic subsequence is "bbbb".


Example 2:
Input:

"cbbd"
Output:
2
One possible longest palindromic subsequence is "bb".


Constraints:

1 <= s.length <= 1000
s consists only of lowercase English letters.
*/

func longestPalindromeSubseq(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return len(s)
	}

	cache := make([][]int, len(s))
	for i := range cache {
		cache[i] = make([]int, len(s))
		cache[i][i] = 1
	}

	return calcPalindromeSubseq(s, 0, len(s)-1, cache)
}

func calcPalindromeSubseq(s string, start, end int, cache [][]int) int {
	if start > end {
		return 0
	}
	if cache[start][end] != 0 {
		return cache[start][end]
	}

	if s[start] == s[end] {
		cache[start][end] = calcPalindromeSubseq(s, start+1, end-1, cache) + 2
	} else {
		v1 := calcPalindromeSubseq(s, start+1, end, cache)
		v2 := calcPalindromeSubseq(s, start, end-1, cache)
		if v1 >= v2 {
			cache[start][end] = v1
		} else {
			cache[start][end] = v2
		}
	}
	return cache[start][end]
}
