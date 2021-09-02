package problem200

import "fmt"

func Judge() {
	grid := [][]byte{[]byte{1, 1, 1, 1, 0}, []byte{1, 1, 0, 1, 0}, []byte{1, 1, 0, 0, 0}, []byte{0, 0, 0, 0, 0}}
	result := numIslands(grid)

	fmt.Println(result)
}

func numIslands(grid [][]byte) int {
	sum := byte(0)

	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			sum += grid[i][j]
		}
	}

	if sum == 0 {
		return 0
	}

	if grid[i][j] == 1 {
		grid[i][j] = 0
		sum
	}

	if i+1 < len(grid[0]) {
		grid = flip(i+1, j, grid)
	}

	if j-1 >= 0 {
		grid = flip(i, j-1, grid)
	}

	if j+1 < len(grid) {
		grid = flip(i, j+1, grid)
	}

	return grid
}
