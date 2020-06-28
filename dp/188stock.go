package main

import (
	"math"
)

/*
Say you have an array for which the i-th element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete at most k transactions.

Note:
You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).

Example 1:

Input: [2,4,1], k = 2
Output: 2
Explanation: Buy on day 1 (price = 2) and sell on day 2 (price = 4), profit = 4-2 = 2.
Example 2:

Input: [3,2,6,5,0,3], k = 2
Output: 7
Explanation: Buy on day 2 (price = 2) and sell on day 3 (price = 6), profit = 6-2 = 4.
             Then buy on day 5 (price = 0) and sell on day 6 (price = 3), profit = 3-0 = 3.

*/

func maxProfit3(k int, prices []int) int {
	if k == 0 {
		return 0
	}
	if k >= len(prices)/2 {
		return maxprofitInf(prices)
	}

	cache := make([][][]int, len(prices))
	for i := range cache {
		cache[i] = make([][]int, k+1)
		for j := range cache[i] {
			cache[i][j] = make([]int, 2)
		}
	}
	for i := range prices {
		cache[i][0][0] = 0
		cache[i][0][1] = math.MinInt64
		for j := k; j >= 1; j-- {
			if i == 0 {
				cache[i][j][0] = 0
				cache[i][j][1] = -prices[i]
			} else {
				cache[i][j][0] = int(math.Max(float64(cache[i-1][j][0]), float64(cache[i-1][j][1]+prices[i])))
				cache[i][j][1] = int(math.Max(float64(cache[i-1][j][1]), float64(cache[i-1][j-1][0]-prices[i])))
			}
		}
	}
	return cache[len(prices)-1][k][0]
}

func maxprofitInf(prices []int) int {
	profit0 := 0
	profit1 := math.MinInt64
	for i := range prices {
		prevProfit0 := profit0
		profit0 = int(math.Max(float64(profit0), float64(profit1+prices[i])))
		profit1 = int(math.Max(float64(profit1), float64(prevProfit0-prices[i])))
	}
	return profit0
}
