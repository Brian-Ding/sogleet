package problems

// NumTrees problem 96
func NumTrees() int {
	return numTrees(6)
}

func numTrees(n int) int {
	if n <= 1 {
		return n
	}
	array := make([]int, 0, 0)
	array = append(array, 1)
	array = append(array, 1)
	for i := 2; i <= n; i++ {
		count := 0
		for j := 0; j < i; j++ {
			count += array[j] * array[i-1-j]
		}
		array = append(array, count)
	}

	return array[n]
}
