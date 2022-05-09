package main

import "fmt"

var sumMatrix [][]int

func minFallingPathSum(matrix [][]int) int {
	if len(matrix) == 0 {
		return 0
	}

	sumMatrix = make([][]int, len(matrix))
	minFallingPathSumDp(matrix, len(matrix)-1)
	minResult := sumMatrix[len(matrix)-1][0]
	for _, v := range sumMatrix[len(matrix)-1] {
		if v < minResult {
			minResult = v
		}
	}
	return minResult
}

func minFallingPathSumDp(matrix [][]int, index int) {
	if len(matrix) == 0 {
		return
	}
	if index < 0 {
		return
	}
	sumMatrix[index] = make([]int, len(matrix[index]))
	if index == 0 {
		copy(sumMatrix[0], matrix[0])
		return
	}

	minFallingPathSumDp(matrix, index-1)
	for i := range matrix[index] {
		sum := matrix[index][i] + sumMatrix[index-1][i]
		if i > 0 && matrix[index][i]+sumMatrix[index-1][i-1] < sum {
			sum = matrix[index][i] + sumMatrix[index-1][i-1]

		}
		if i < len(matrix[index])-1 && matrix[index][i]+sumMatrix[index-1][i+1] < sum {
			sum = matrix[index][i] + sumMatrix[index-1][i+1]
		}
		fmt.Printf("i: %d, index: %d, matrix[index][i]: %d, sumMatrix[index-1]: %v\n",
			i, index, matrix[index][i], sumMatrix[index-1])
		sumMatrix[index][i] = sum
	}
	fmt.Printf("sumMatrix: %v\n", sumMatrix)
	// [[100,-42,-46,-41],[31,97,10,-10],[-58,-51,82,89],[51,81,69,-51]]
}
