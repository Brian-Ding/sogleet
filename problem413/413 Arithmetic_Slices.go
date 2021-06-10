package problems

import (
	"math"
)

// NumberOfArithmeticSlices problem 413
func NumberOfArithmeticSlices(A []int) int {
	return numberOfArithmeticSlices(A)
}

func numberOfArithmeticSlices(A []int) int {
	row := len(A)
	array := newArray(row, row, math.MaxInt32)
	count := 0

	for i := 0; i < row; i++ {
		for j := i + 2; j < row; j++ {
			if j == i+2 {
				if A[j]-A[j-1] == A[j-1]-A[i] {
					array[i][j] = A[j] - A[j-1]
				}
			} else {
				if A[j]-A[j-1] == array[i][j-1] {
					array[i][j] = array[i][j-1]
				}
			}

			if array[i][j] != math.MaxInt32 {
				count++
			}

		}
	}

	return count
}
