package main

/*
Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.

'.' Matches any single character.
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).

Note:

s could be empty and contains only lowercase letters a-z.
p could be empty and contains only lowercase letters a-z, and characters like . or *.
Example 1:

Input:
s = "aa"
p = "a"
Output: false
Explanation: "a" does not isMatch the entire string "aa".
Example 2:

Input:
s = "aa"
p = "a*"
Output: true
Explanation: '*' means zero or more of the preceding element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
Example 3:

Input:
s = "ab"
p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
Example 4:

Input:
s = "aab"
p = "c*a*b"
Output: true
Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore, it matches "aab".
Example 5:

Input:
s = "mississippi"
p = "mis*is*p*."
Output: false
*/

const (
	UNMATCHED = iota
	MATCH
	NOTMATCH
)

var matchCache = make(map[string]map[string]int)

// isMatch
// time: O(m*n)
// storage: O(m*n)
func isMatch(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	if len(p) == 0 {
		return false
	}
	if p[0] == '*' {
		return false
	}

	if v1, ok := matchCache[s]; ok {
		if v2, ok := v1[p]; ok {
			if v2 == MATCH {
				return true
			} else if v2 == NOTMATCH {
				return false
			}
		}
	} else {
		matchCache[s] = make(map[string]int)
	}

	firstMatch := false
	if len(s) > 0 {
		firstMatch = s[0] == p[0] || p[0] == '.'
	}

	result := false
	if firstMatch {
		if len(p) > 1 && p[1] == '*' {
			result = isMatch(s[1:], p) || isMatch(s[1:], p[2:]) || isMatch(s, p[2:])
		} else {
			result = isMatch(s[1:], p[1:])
		}
	} else {
		if len(p) > 1 && p[1] == '*' {
			result = isMatch(s, p[2:])
		}
	}
	if result {
		matchCache[s][p] = MATCH
	} else {
		matchCache[s][p] = NOTMATCH
	}
	return result
}
