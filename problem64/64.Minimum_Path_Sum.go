package problems

// MinPathSum problem 64
func MinPathSum(grid [][]int) int {
	return minPathSum(grid)
}

func minPathSum(grid [][]int) int {
	row := len(grid)
	column := len(grid[0])
	array := newArray(row, column, 0)
	array[0][0] = grid[0][0]
	for i := 1; i < row; i++ {
		array[i][0] = grid[i][0] + array[i-1][0]
	}
	for j := 1; j < column; j++ {
		array[0][j] = grid[0][j] + array[0][j-1]
	}

	for i := 1; i < row; i++ {
		for j := 1; j < column; j++ {
			array[i][j] = min(array[i-1][j], array[i][j-1]) + grid[i][j]
		}
	}

	return array[row-1][column-1]
}
