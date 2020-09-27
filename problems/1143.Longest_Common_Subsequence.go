package problems

// LongestCommonSubsequence problem 1143
func LongestCommonSubsequence(text1 string, text2 string) int {
	return longestCommonSubsequence(text1, text2)
}

func longestCommonSubsequence(text1 string, text2 string) int {
	row := len(text1) + 1
	column := len(text2) + 1
	array := newArray(row, column, 0)

	for i := 1; i < row; i++ {
		for j := 1; j < column; j++ {
			if text1[i-1] == text2[j-1] {
				array[i][j] = array[i-1][j-1] + 1
			} else {
				array[i][j] = max(array[i][j-1], array[i-1][j])
			}
		}
	}

	return array[row-1][column-1]
}
