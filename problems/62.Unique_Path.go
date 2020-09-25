package problems

func uniquePaths(m int, n int) int {
	// row: m
	// column: n
	if m == 1 || n == 1 {
		return 1
	}

	robotMap := make([][]int, m)
	for i := range robotMap {
		robotMap[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				robotMap[i][j] = 1
			} else {
				robotMap[i][j] = robotMap[i-1][j] + robotMap[i][j-1]
			}
		}
	}

	return robotMap[m-1][n-1]
}
