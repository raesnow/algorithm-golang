package main

/*
Given a 2D board containing 'X' and 'O' (the letter O), capture all regions surrounded by 'X'.

A region is captured by flipping all 'O's into 'X's in that surrounded region.

Example:

X X X X
X O O X
X X O X
X O X X
After running your function, the board should be:

X X X X
X X X X
X X X X
X O X X
Explanation:

Surrounded regions shouldn’t be on the border, which means that any 'O' on the border of the board are not flipped to 'X'. Any 'O' that is not on the border and it is not connected to an 'O' on the border will be flipped to 'X'. Two cells are connected if they are adjacent cells connected horizontally or vertically.
*/

func solve(board [][]byte) {
	if len(board) == 0 {
		return
	}
	checked := make([]int, len(board)*len(board[0]))

	for i := range board {
		for j := range board[i] {
			if i == 0 || i == len(board)-1 || j == 0 || j == len(board[0])-1 {
				if board[i][j] == 'O' {
					checkSurounded(board, i, j, checked)
					continue
				}
			}
			if board[i][j] == 'X' {
				checked[i*len(board[0])+j] = X
			}
		}
	}

	for i := range board {
		for j := range board[i] {
			if checked[i*len(board[0])+j] == NONE {
				board[i][j] = 'X'
			}
		}
	}

}

const (
	NONE = iota
	X
	O
)

func checkSurounded(board [][]byte, i, j int, checked []int) {
	if i >= len(board) || i < 0 || j >= len(board[0]) || j < 0 {
		return
	}
	if checked[i*len(board[0])+j] != NONE {
		return
	}

	checked[i*len(board[0])+j] = O

	if i > 0 && board[i-1][j] == 'O' {
		checkSurounded(board, i-1, j, checked)
	}
	if i < len(board)-1 && board[i+1][j] == 'O' {
		checkSurounded(board, i+1, j, checked)
	}
	if j > 0 && board[i][j-1] == 'O' {
		checkSurounded(board, i, j-1, checked)
	}
	if j < len(board[0])-1 && board[i][j+1] == 'O' {
		checkSurounded(board, i, j+1, checked)
	}
}

type UnionFind struct {
	count int
	root  []int
	size  []int
}

func createUnionFind(count int) *UnionFind {
	uf := &UnionFind{
		count: count,
		root:  make([]int, count),
		size:  make([]int, count),
	}
	for i := range uf.root {
		uf.root[i] = i
	}
	for i := range uf.size {
		uf.size[i] = 1
	}
	return uf
}

func (uf *UnionFind) union(i, j int) {
	rootI := uf.find(i)
	rootJ := uf.find(j)

	if rootI != rootJ {
		if uf.size[rootI] >= uf.size[rootJ] {
			uf.root[rootJ] = rootI
			uf.size[rootI] += uf.size[rootJ]
		} else {
			uf.root[rootI] = rootJ
			uf.size[rootJ] += uf.size[rootI]
		}
		uf.count--
	}
}

func (uf *UnionFind) find(i int) int {
	for uf.root[i] != i {
		uf.root[i] = uf.root[uf.root[i]]
		i = uf.root[i]
	}
	return i
}

func (uf *UnionFind) connected(i, j int) bool {
	return uf.find(i) == uf.find(j)
}

func solve1(board [][]byte) {
	if len(board) == 0 {
		return
	}

	uf := createUnionFind(len(board)*len(board[0]) + 1)
	for i := range board {
		for j := range board[i] {
			if i == 0 || i == len(board)-1 || j == 0 || j == len(board[0])-1 {
				if board[i][j] == 'O' {
					uf.union(i*len(board[0])+j, len(board)*len(board[0]))
				}
			} else if board[i][j] == 'O' {
				directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
				for _, values := range directions {
					if board[i+values[0]][j+values[1]] == 'O' {
						uf.union((i+values[0])*len(board[0])+j+values[1], i*len(board[0])+j)
					}
				}
			}
		}
	}
	for i := range board {
		for j := range board[i] {
			if board[i][j] == 'O' && !uf.connected(i*len(board[0])+j, len(board)*len(board[0])) {
				board[i][j] = 'X'
			}
		}
	}
}

func main() {
	solve1([][]byte{{'X', 'X', 'X', 'X'}, {'X', 'O', 'O', 'X'}, {'X', 'X', 'O', 'X'}, {'X', 'O', 'X', 'X'}})
}
