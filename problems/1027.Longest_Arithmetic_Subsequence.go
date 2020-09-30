package problems

func LongestArithSeqLength() int {
	A := []int{9, 4, 7, 2, 10}
	return longestArithSeqLength(A)
}

func longestArithSeqLength(A []int) int {
	array := make([]map[int]int, 0, len(A))
	array = append(array, map[int]int{A[0]: 1})

	for i := 1; i < len(A); i++ {
		for j := 1; j >= 0; j++ {
			dict := array[i-j]
			for k, v := range dict {

			}
		}
	}
}
