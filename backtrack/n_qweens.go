package main

var resultQween [][][]int

func nQweens(n int) [][][]int {
	if n == 0 {
		return nil
	}
	if n == 1 {
		return [][][]int{{{1}}}
	}

	path := make([][]int, n)
	for i := range path {
		path[i] = make([]int, n)
	}
	backtrackQween(path, 0, n)
	return resultQween
}

func backtrackQween(path [][]int, calculated, endLength int) {
	if calculated == endLength {
		result := make([][]int, endLength)
		for i := range result {
			result[i] = make([]int, endLength)
			for j := range result[i] {
				result[i][j] = path[i][j]
			}
		}
		resultQween = append(resultQween, result)
		return
	}

	for i := 0; i < endLength; i++ {
		if checkOk(calculated, i, path) {
			path[calculated][i] = 1
			backtrackQween(path, calculated+1, endLength)
			path[calculated][i] = 0
		}
	}

	return
}

func checkOk(x, y int, matrix [][]int) bool {
	for i := x - 1; i >= 0; i-- {
		if matrix[i][y] == 1 {
			return false
		}
	}

	for i, j := x-1, y-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if matrix[i][j] == 1 {
			return false
		}
	}

	for i, j := x-1, y+1; i >= 0 && j < len(matrix); i, j = i-1, j+1 {
		if matrix[i][j] == 1 {
			return false
		}
	}

	return true
}
