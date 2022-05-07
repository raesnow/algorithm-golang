package array

type NumMatrix struct {
	matrix  [][]int
	preSums [][]int
}

func ConstructorMatrix(matrix [][]int) NumMatrix {
	preSums := make([][]int, len(matrix))
	for i, row := range matrix {
		preSums[i] = make([]int, len(row))
		for j, value := range row {
			if i == 0 && j == 0 {
				preSums[i][j] = value
			} else if i == 0 {
				preSums[i][j] = preSums[i][j-1] + value
			} else if j == 0 {
				preSums[i][j] = preSums[i-1][j] + value
			} else {
				preSums[i][j] = preSums[i-1][j] + preSums[i][j-1] - preSums[i-1][j-1] + value
			}
		}
	}
	return NumMatrix{
		matrix:  matrix,
		preSums: preSums,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	if row1 > row2 || col1 > col2 || row1 < 0 || col1 < 0 || len(this.matrix) == 0 {
		return 0
	}
	if row2 >= len(this.matrix) {
		row2 = len(this.matrix) - 1
	}
	if col2 >= len(this.matrix[0]) {
		col2 = len(this.matrix[0]) - 1
	}

	if col1 == 0 && row1 == 0 {
		return this.preSums[row2][col2]
	}
	if col1 == 0 {
		return this.preSums[row2][col2] - this.preSums[row1-1][col2]
	}
	if row1 == 0 {
		return this.preSums[row2][col2] - this.preSums[row2][col1-1]
	}
	return this.preSums[row2][col2] - this.preSums[row2][col1-1] - this.preSums[row1-1][col2] + this.preSums[row1-1][col1-1]
}
