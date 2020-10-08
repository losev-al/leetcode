package main

func uniquePaths(m int, n int) int {
	data := make([][]int,m)
	for i := 0; i < m; i++ {
		data[i] = make([]int, n)
		data[i][0] = 1
	}
	for i := 0; i < n; i++ {
		data[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			data[i][j] = data[i-1][j] + data[i][j-1]
		}
	}
	return data[m-1][n-1]
}



//func uniquePaths(m int, n int) int {
//	count := 0
//	buildPath(1, 1, m, n, &count)
//	return count
//}
//
//func buildPath(r int, c int, m int, n int, count *int) {
//	if r == m && c == n {
//		*count += 1
//		return
//	}
//	if r < m {
//		buildPath(r+1, c, m, n, count)
//	}
//	if c < n {
//		buildPath(r, c+1, m, n, count)
//	}
//}
