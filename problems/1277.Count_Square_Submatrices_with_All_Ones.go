package problems

import (
	"math"
)

func CountSquares(matrix [][]int) int {
	return countSquares(matrix)
}

func countSquares(matrix [][]int) int {
	row := len(matrix)
	column := len(matrix[0])

	if row == 0 || column == 0 {
		return 0
	}

	count := 0
	result := make([][]int, 0, row+1)
	for i := 0; i < row+1; i++ {
		result = append(result, make([]int, 0, column+1))
		for j := 0; j < column+1; j++ {
			result[i] = append(result[i], 0)
		}
	}

	for i := 1; i < row+1; i++ {
		for j := 1; j < column+1; j++ {
			if matrix[i-1][j-1] == 0 {
				result[i][j] = 0
			} else {
				result[i][j] = int(math.Min(math.Min(float64(result[i-1][j]), float64(result[i][j-1])), float64(result[i-1][j-1]))) + matrix[i-1][j-1]
				count += result[i][j]
			}
		}
	}

	return count
}
