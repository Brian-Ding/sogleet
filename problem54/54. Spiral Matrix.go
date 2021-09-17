package problem54

import "fmt"

func Judge() {
	// matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	matrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	fmt.Println(spiralOrder(matrix))
}

func spiralOrder(matrix [][]int) []int {
	rowN := len(matrix)
	columnN := len(matrix[0])

	maps := make([][]bool, rowN)
	for i := 0; i < rowN; i++ {
		maps[i] = make([]bool, columnN)
	}
	result := make([]int, rowN*columnN)

	row := 0
	column := 0
	for i := 0; i < len(result); i++ {
		result[i] = matrix[row][column]
		maps[row][column] = true

		right := column != columnN-1 && !maps[row][column+1]
		down := row != rowN-1 && !maps[row+1][column]
		left := column != 0 && !maps[row][column-1]
		up := row != 0 && !maps[row-1][column]

		if right && down {
			column++
			continue
		}

		if down && left {
			row++
			continue
		}

		if left && up {
			column--
			continue
		}

		if up && right {
			row--
			continue
		}

		if right {
			column++
			continue
		}

		if down {
			row++
			continue
		}

		if left {
			column--
			continue
		}

		if up {
			row--
			continue
		}

		break
	}

	return result
}
