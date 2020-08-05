package calculator

import "fmt"

/*
Implement a basic calculator to evaluate a simple expression string.

The expression string may contain open ( and closing parentheses ), the plus + or minus sign -, non-negative integers and empty spaces .

Example 1:

Input: "1 + 1"
Output: 2
Example 2:

Input: " 2-1 + 2 "
Output: 3
Example 3:

Input: "(1+(4+5+2)-3)+(6+8)"
Output: 23
Note:
You may assume that the given expression is always valid.
Do not use the eval built-in library function.
*/

func calculate(s string) int {
	result, end := calculateHelper(s, 0)
	if end < len(s) {
		fmt.Printf("calculate stop at index: %d", end-1)
	}
	return result
}

func calculateHelper(s string, startIndex int) (result, endIndex int) {
	stack := make([]int, 0)
	num := 0
	sign := 1
	i := startIndex
	for ; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		} else if c == ')' {
			break
		} else {
			switch c {
			case '+':
				stack = append(stack, num*sign)
				sign = 1
				num = 0
			case '-':
				stack = append(stack, num*sign)
				sign = -1
				num = 0
			case '(':
				num, i = calculateHelper(s, i+1)
				stack = append(stack, num*sign)
				sign = 1
				num = 0
			case ' ':
				break
			}
		}
	}
	stack = append(stack, num*sign)
	sum := 0
	for _, v := range stack {
		sum += v
	}
	return sum, i
}
