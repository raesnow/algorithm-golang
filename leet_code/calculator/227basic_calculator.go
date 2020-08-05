package calculator

import "fmt"

/*
Implement a basic calculator to evaluate a simple expression string.

The expression string contains only non-negative integers, +, -, *, / operators and empty spaces . The integer division should truncate toward zero.

Example 1:

Input: "3+2*2"
Output: 7
Example 2:

Input: " 3/2 "
Output: 1
Example 3:

Input: " 3+5 / 2 "
Output: 5
Note:

You may assume that the given expression is always valid.
Do not use the eval built-in library function.
*/

func calculate1(s string) int {
	result, end := calculateHelper1(s, 0)
	if end < len(s) {
		fmt.Printf("calculate stop at index: %d", end-1)
	}
	return result
}

func calculateHelper1(s string, startIndex int) (result, endIndex int) {
	stack := make([]int, 0)
	num := 0
	var sign uint8 = '+'
	i := startIndex
	for ; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		} else if c == ')' {
			break
		} else if c == '+' || c == '-' || c == '*' || c == '/' || c == '(' {
			if c == '(' {
				num, i = calculateHelper1(s, i+1)
				c = '+'
			}
			switch sign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] = stack[len(stack)-1] * num
			case '/':
				stack[len(stack)-1] = stack[len(stack)-1] / num
			}
			sign = c
			num = 0
		} else if c != ' ' {
			fmt.Println("not recognize: ", c)
		}
	}

	switch sign {
	case '+':
		stack = append(stack, num)
	case '-':
		stack = append(stack, -num)
	case '*':
		stack[len(stack)-1] = stack[len(stack)-1] * num
	case '/':
		stack[len(stack)-1] = stack[len(stack)-1] / num
	}
	sum := 0
	for _, v := range stack {
		sum += v
	}
	return sum, i
}
