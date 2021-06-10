package problems

// LongestPalindromeSubseq problem 516
func LongestPalindromeSubseq() int {
	a := "bbbab"
	return longestPalindromeSubseq(a)
}

func longestPalindromeSubseq(s string) int {
	width := len(s)
	array := newArray(width, width, 0)
	array[0][0] = 1
	for i := 1; i < width; i++ {
		array[i][i] = 1
		if s[i-1] == s[i] {
			array[i-1][i] = 2
		} else {
			array[i-1][i] = 1
		}
	}

	for i := 2; i < width; i++ {
		for j := 0; j < width && j+i < width; j++ {
			if s[j] == s[j+i] {
				array[j][j+i] = array[j+1][j+i-1] + 2
			} else {
				array[j][j+i] = max(array[j][j+i-1], array[j+1][j+i])
			}
		}
	}

	return array[0][width-1]
}
