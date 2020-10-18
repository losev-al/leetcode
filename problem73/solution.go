package main

func setZeroes(matrix [][]int) {
	row := len(matrix)
	if row == 0 {
		return
	}
	col := len(matrix[0])
	if col == 0 {
		return
	}
	zerroCols := make(map[int]bool)
	zerroRows := make(map[int]bool)
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			if matrix[r][c] == 0 {
				zerroCols[c] = true
				zerroRows[r] = true
			}
		}
	}
	for r, _ := range zerroRows {
		for i := 0; i < col; i++ {
			matrix[r][i] = 0
		}
	}
	for c, _ := range zerroCols {
		for i := 0; i < row; i++ {
			matrix[i][c] = 0
		}
	}
}
