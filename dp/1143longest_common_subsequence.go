package main

/*
Given two strings text1 and text2, return the length of their longest common subsequence.

A subsequence of a string is a new string generated from the original string with some characters
(can be none) deleted without changing the relative order of the remaining characters.
(eg, "ace" is a subsequence of "abcde" while "aec" is not).
A common subsequence of two strings is a subsequence that is common to both strings.

If there is no common subsequence, return 0.

Example 1:

Input: text1 = "abcde", text2 = "ace"
Output: 3
Explanation: The longest common subsequence is "ace" and its length is 3.

Example 2:

Input: text1 = "abc", text2 = "abc"
Output: 3
Explanation: The longest common subsequence is "abc" and its length is 3.

Example 3:

Input: text1 = "abc", text2 = "def"
Output: 0
Explanation: There is no such common subsequence, so the result is 0.


Constraints:

1 <= text1.length <= 1000
1 <= text2.length <= 1000
The input strings consist of lowercase English characters only.

*/

func longestCommonSubsequence(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}

	memo := make([][]int, len(text1)+1)
	for i := range memo {
		memo[i] = make([]int, len(text2)+1)
	}

	for i := 1; i <= len(text1); i++ {
		for j := 1; j <= len(text2); j++ {
			if text1[i-1] == text2[j-1] {
				memo[i][j] = memo[i-1][j-1] + 1
			} else {
				memo[i][j] = max(memo[i-1][j], memo[i][j-1])
			}
		}
	}
	return memo[len(text1)][len(text2)]
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

var memo [][]int

func longestCommonSubsequence1(text1 string, text2 string) int {
	if len(text1) == 0 || len(text2) == 0 {
		return 0
	}

	memo = make([][]int, len(text1)+1)
	for i := range memo {
		memo[i] = make([]int, len(text2)+1)
	}

	return dpLCS(text1, text2, len(text1), len(text2))
}

func dpLCS(text1, text2 string, i, j int) int {
	if i == 0 || j == 0 {
		return 0
	}
	if memo[i][j] != 0 {
		return memo[i][j]
	}

	if text1[i-1] == text2[j-1] {
		memo[i][j] = dpLCS(text1, text2, i-1, j-1) + 1
	} else {
		memo[i][j] = max(dpLCS(text1, text2, i-1, j), dpLCS(text1, text2, i, j-1))
	}
	return memo[i][j]
}
