package main

/*
Given two words word1 and word2, find the minimum number of operations required to
convert word1 to word2.

You have the following 3 operations permitted on a word:

Insert a character
Delete a character
Replace a character

Example 1:

Input: word1 = "horse", word2 = "ros"
Output: 3
Explanation:
horse -> rorse (replace 'h' with 'r')
rorse -> rose (remove 'r')
rose -> ros (remove 'e')

Example 2:

Input: word1 = "intention", word2 = "execution"
Output: 5
Explanation:
intention -> inention (remove 't')
inention -> enention (replace 'i' with 'e')
enention -> exention (replace 'n' with 'x')
exention -> exection (replace 'n' with 'c')
exection -> execution (insert 'u')
*/

var cache [][]int

func minDistance(word1 string, word2 string) int {
	cache = make([][]int, len(word1))
	for i := range cache {
		cache[i] = make([]int, len(word2))
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	return findDistance(word1, word2, len(word1)-1, len(word2)-1)
}

func findDistance(word1 string, word2 string, i int, j int) int {
	if i == -1 {
		return j + 1
	}
	if j == -1 {
		return i + 1
	}

	if cache[i][j] != -1 {
		return cache[i][j]
	}

	result := findDistance(word1, word2, i-1, j-1)
	if word1[i] != word2[j] {
		l1 := findDistance(word1, word2, i-1, j) + 1
		l2 := findDistance(word1, word2, i, j-1) + 1
		l3 := findDistance(word1, word2, i-1, j-1) + 1
		result = min(l1, l2, l3)
	}
	cache[i][j] = result
	return result
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

func minDistance1(word1 string, word2 string) int {
	if len(word1) == 0 {
		return len(word2)
	}
	if len(word2) == 0 {
		return len(word1)
	}

	cache := make([][]int, len(word1)+1)
	for i := range cache {
		cache[i] = make([]int, len(word2)+1)
	}

	for i := 0; i <= len(word1); i++ {
		cache[i][0] = i
	}
	for j := 0; j <= len(word2); j++ {
		cache[0][j] = j
	}

	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			result := cache[i-1][j-1]
			if word1[i-1] != word2[j-1] {
				l1 := cache[i-1][j] + 1
				l2 := cache[i][j-1] + 1
				l3 := cache[i-1][j-1] + 1
				result = min(l1, l2, l3)
			}
			cache[i][j] = result
		}
	}
	return cache[len(word1)][len(word2)]
}
