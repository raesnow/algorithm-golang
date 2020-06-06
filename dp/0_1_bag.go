package main

func bag(weight int, values, weights []int) int {
	if len(values) == 0 || len(weights) == 0 || len(values) != len(weights) {
		return 0
	}

	dp := make([][]int, len(values)+1)
	for i := range dp {
		dp[i] = make([]int, weight+1)
	}

	for i := 1; i <= len(values); i++ {
		for j := 1; j <= weight; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= weights[i-1] {
				newWeight := dp[i-1][j-weights[i-1]] + values[i-1]
				if newWeight > dp[i-1][j] {
					dp[i][j] = newWeight
				}
			}
		}
	}
	return dp[len(values)][weight]
}
