package problems

import (
	"math"
)

// MinFallingPathSum problem 931
func MinFallingPathSum(A [][]int) int {
	return minFallingPathSum(A)
}

func minFallingPathSum(A [][]int) int {
	row := len(A)
	column := len(A[0])

	result := newArray(row, column, math.MaxInt32)
	for j := 0; j < column; j++ {
		result[row-1][j] = A[row-1][j]
	}

	for i := row - 2; i >= 0; i-- {
		for j := column - 1; j >= 0; j-- {
			temp := math.MaxInt32
			if j > 0 {
				temp = min(temp, result[i+1][j-1])
			}
			temp = min(temp, result[i+1][j])
			if j < column-1 {
				temp = min(temp, result[i+1][j+1])
			}

			result[i][j] = temp + A[i][j]
		}
	}

	answer := result[0][0]
	for j := 1; j < column; j++ {
		answer = min(answer, result[0][j])
	}

	return answer
}
