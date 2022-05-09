package main

import (
	"math"
)

var coinChangeMemo = make(map[int]int)

func coinChange2(coins []int, amount int) int {
	coinChangeMemo = make(map[int]int)
	return coinChangeDP2(coins, amount)
}

func coinChangeDP2(coins []int, amount int) int {
	if amount < 0 {
		return -1
	}
	if amount == 0 {
		return 0
	}
	val, ok := coinChangeMemo[amount]
	if ok {
		return val
	}

	result := math.MaxInt
	for _, coin := range coins {
		sub := coinChangeDP2(coins, amount-coin)
		if sub == -1 {
			continue
		}
		if sub+1 < result {
			result = sub + 1
		}
	}
	if result == math.MaxInt {
		coinChangeMemo[amount] = -1
	} else {
		coinChangeMemo[amount] = result
	}
	return coinChangeMemo[amount]
}
