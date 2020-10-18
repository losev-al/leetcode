package main

func exist(board [][]byte, word string) bool {
	row := len(board)
	if row == 0 {
		return false
	}
	col := len(board[0])
	if col == 0 {
		return false
	}
	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
	}
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			result := find(board, word, &visited, r, c, row, col, 0)
			if result {
				return true
			}
		}
	}
	return false
}

func find(board [][]byte, word string, visited *[][]bool, r int, c int, row int, col int, index int) bool {
	if r >= row || r < 0 || c >= col || c < 0 || (*visited)[r][c] {
		return false
	}
	if board[r][c] == word[index] {
		(*visited)[r][c] = true
		index++
		if index == len(word) {
			return true
		}
		if find(board, word, visited, r+1, c, row, col, index) {
			return true
		}
		if find(board, word, visited, r-1, c, row, col, index) {
			return true
		}
		if find(board, word, visited, r, c+1, row, col, index) {
			return true
		}
		if find(board, word, visited, r, c-1, row, col, index) {
			return true
		}
		(*visited)[r][c] = false
	}
	return false
}
