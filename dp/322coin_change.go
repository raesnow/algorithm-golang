package main

/*
You are given coins of different denominations and a total amount of money amount. Write a function to compute the fewest
number of coins that you need to make up that amount. If that amount of money cannot be made up by any combination of the coins, return -1.

Example 1:

Input: coins = [1, 2, 5], amount = 11
Output: 3
Explanation: 11 = 5 + 5 + 1
Example 2:

Input: coins = [2], amount = 3
Output: -1
Note:
You may assume that you have an infinite number of each kind of coin.

*/

func coinChange(coins []int, amount int) int {
	cache := make(map[int]int)
	return coinChangeDP(coins, amount, cache)
}

func coinChangeDP(coins []int, amount int, cache map[int]int) int {
	if amount == 0 {
		return 0
	}
	if len(coins) == 0 {
		return -1
	}
	if v, ok := cache[amount]; ok {
		return v
	}

	result := -1
	for _, coin := range coins {
		if amount < coin {
			continue
		}
		others := coinChange(coins, amount-coin)
		if others == -1 {
			continue
		}

		if result == -1 {
			result = others + 1
		} else if result > others+1 {
			result = others + 1
		}
	}
	cache[amount] = result
	return result
}

func coinChange1(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	cache := make([]int, amount+1)
	cache[0] = 0

	for i := 1; i <= amount; i++ {
		cache[i] = -1
		for _, coin := range coins {
			if i < coin {
				continue
			}
			if cache[i-coin] == -1 {
				continue
			}
			if cache[i] == -1 {
				cache[i] = cache[i-coin] + 1
			} else if cache[i] > cache[i-coin]+1 {
				cache[i] = cache[i-coin] + 1
			}
		}
	}
	return cache[amount]
}
