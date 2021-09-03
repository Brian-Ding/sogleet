package problem200

import "fmt"

func Judge() {
	//grid := [][]byte{{1, 1, 1, 1, 0}, {1, 1, 0, 1, 0}, {1, 1, 0, 0, 0}, {0, 0, 0, 0, 0}}
	//grid := [][]byte{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 1, 1}}
	grid := [][]byte{{1, 1, 0, 0, 0}, {1, 1, 0, 0, 0}, {0, 0, 1, 0, 0}, {0, 0, 0, 1, 1}}
	result := numIslands(grid)

	fmt.Println(result)
}

func numIslands(grid [][]byte) int {
	sum := byte(0)
	count := 0

	rowSum := make([]byte, len(grid))

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			sum += grid[i][j]
			rowSum[i] += grid[i][j]
		}
	}

	if sum == 0 {
		return 0
	}

	for i := 0; i < len(grid); i++ {
		if rowSum[i] == 0 {
			continue
		}

		for j := 0; j < len(grid[i]); j++ {
			if rowSum[i] == 0 {
				break
			}

			if flip(i, j, &grid, &rowSum, &sum) {
				count++
			}
		}
	}

	return count
}

func flip(i int, j int, grid *[][]byte, rowSum *[]byte, sum *byte) bool {

	if (*grid)[i][j] == 0 {
		return false
	}

	(*grid)[i][j] = 0
	(*rowSum)[i]--
	*sum--

	if i >= 1 {
		flip(i-1, j, grid, rowSum, sum)
	}

	if i+1 < len(*grid) {
		flip(i+1, j, grid, rowSum, sum)
	}

	if j+1 < len((*grid)[i]) {
		flip(i, j+1, grid, rowSum, sum)
	}

	return true
}
