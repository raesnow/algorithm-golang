package main

var fibMemo = map[int]int{
	0: 0,
	1: 1,
}

func fib2(n int) int {
	if n == 0 || fibMemo[n] != 0 {
		return fibMemo[n]
	}

	result := fib2(n-1) + fib2(n-2)
	fibMemo[n] = result
	return result
}
