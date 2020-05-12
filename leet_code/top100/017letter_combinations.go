package main

/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
Note:

Although the above answer is in lexicographical order, your answer could be in any order you want.
*/

var letterMap = map[byte][]string{
	byte('2'): []string{"a", "b", "c"},
	byte('3'): []string{"d", "e", "f"},
	byte('4'): []string{"g", "h", "i"},
	byte('5'): []string{"j", "k", "l"},
	byte('6'): []string{"m", "n", "o"},
	byte('7'): []string{"p", "q", "r", "s"},
	byte('8'): []string{"t", "u", "v"},
	byte('9'): []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	if len(digits) == 1 {
		return letterMap[digits[0]]
	}

	others := letterCombinations(digits[1:])
	result := make([]string, 0)
	for _, first := range letterMap[digits[0]] {
		for _, other := range others {
			result = append(result, first+other)
		}
	}
	return result
}
