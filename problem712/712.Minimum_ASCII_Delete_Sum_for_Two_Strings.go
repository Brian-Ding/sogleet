package problems

import "github.com/Brian-Ding/sogleet/common"

// MinimumDeleteSum problem 712
func MinimumDeleteSum(s1 string, s2 string) int {
	return minimumDeleteSum(s1, s2)
}

func minimumDeleteSum(s1 string, s2 string) int {
	row := len(s1) + 1
	column := len(s2) + 1
	array := common.NewArray(row, column, 0)

	for i := 1; i < row; i++ {
		array[i][0] = array[i-1][0] + int(s1[i-1])
	}

	for j := 1; j < column; j++ {
		array[0][j] = array[0][j-1] + int(s2[j-1])
	}

	for i := 1; i < row; i++ {
		for j := 1; j < column; j++ {
			if s1[i-1] == s2[j-1] {
				array[i][j] = array[i-1][j-1]
			} else {
				array[i][j] = common.Min(array[i][j-1]+int(s2[j-1]), array[i-1][j]+int(s1[i-1]))
			}
		}
	}

	return array[row-1][column-1]
}
