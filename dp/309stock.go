package main

import "math"

/*
Say you have an array for which the ith element is the price of a given stock on day i.

Design an algorithm to find the maximum profit. You may complete as many transactions as you like
(ie, buy one and sell one share of the stock multiple times) with the following restrictions:

You may not engage in multiple transactions at the same time (ie, you must sell the stock before you buy again).
After you sell your stock, you cannot buy stock on next day. (ie, cooldown 1 day)
Example:

Input: [1,2,3,0,2]
Output: 3
Explanation: transactions = [buy, sell, cooldown, buy, sell]

*/

func maxProfit4(prices []int) int {
	prevProfit0 := 0
	profit0 := 0
	profit1 := math.MinInt64
	for i := range prices {
		tmp := profit0
		profit0 = int(math.Max(float64(profit0), float64(profit1+prices[i])))
		profit1 = int(math.Max(float64(profit1), float64(prevProfit0-prices[i])))
		prevProfit0 = tmp

	}
	return profit0
}
