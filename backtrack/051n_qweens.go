package main

/*
The n-queens puzzle is the problem of placing n queens on an n×n chessboard such that no two queens attack each other.

Given an integer n, return all distinct solutions to the n-queens puzzle.

Each solution contains a distinct board configuration of the n-queens' placement, where 'Q' and '.' both indicate a
queen and an empty space respectively.

Example:

Input: 4
Output: [
 [".Q..",  // Solution 1
  "...Q",
  "Q...",
  "..Q."],

 ["..Q.",  // Solution 2
  "Q...",
  "...Q",
  ".Q.."]
]
Explanation: There exist two distinct solutions to the 4-queens puzzle as shown above.

*/

var resultQueen [][]string

func solveNQueens(n int) [][]string {
	if n == 0 {
		return nil
	}
	if n == 1 {
		return [][]string{{"Q"}}
	}
	resultQueen = make([][]string, 0)

	path := make([]string, n)
	initValues := make([]byte, n)
	for i := range initValues {
		initValues[i] = '.'
	}
	for i := range path {
		now := make([]byte, n)
		copy(now, initValues)
		path[i] = string(now)
	}

	backtrackQween(path, 0, n)
	return resultQueen
}

func backtrackQween(path []string, calculated, endLength int) {
	if calculated == endLength {
		result := make([]string, endLength)
		copy(result, path)
		resultQueen = append(resultQueen, result)
		return
	}

	for i := 0; i < endLength; i++ {
		if checkOk(calculated, i, path) {
			initValues := make([]byte, len(path))
			for j := range initValues {
				initValues[j] = '.'
			}
			initValues[i] = 'Q'
			path[calculated] = string(initValues)
			backtrackQween(path, calculated+1, endLength)
			initValues[i] = '.'
			path[calculated] = string(initValues)
		}
	}

	return
}

func checkOk(x, y int, matrix []string) bool {
	for i := x - 1; i >= 0; i-- {
		if matrix[i][y] == 'Q' {
			return false
		}
	}

	for i, j := x-1, y-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if matrix[i][j] == 'Q' {
			return false
		}
	}

	for i, j := x-1, y+1; i >= 0 && j < len(matrix); i, j = i-1, j+1 {
		if matrix[i][j] == 'Q' {
			return false
		}
	}

	return true
}
