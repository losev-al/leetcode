package main

func spiralOrder(matrix [][]int) []int {
	m := len(matrix)
	if m == 0 {
		return make([]int, 0)
	}
	n := len(matrix[0])
	c := 0
	r := 0
	dC := 1
	dR := 0
	result := make([]int, 0, n*m)
	for {
		result = append(result, matrix[r][c])
		if len(result) == len(matrix)*len(matrix[0]) {
			break
		}
		if c == n-1 && dC == 1 {
			dC = 0
			dR = 1
		} else if r == m-1 && dR == 1 {
			dR = 0
			dC = -1
		} else if c == len(matrix[0])-n && dC == -1 {
			dC = 0
			dR = -1
		} else if r == len(matrix)-m+1 && dR == -1 {
			dC = 1
			dR = 0
			n--
			m--
		}
		r += dR
		c += dC
	}
	return result
}
