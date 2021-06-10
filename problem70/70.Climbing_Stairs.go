package problems

// ClimbStairs problem 70
func ClimbStairs(n int) int {
	return climbStairs(n)
}

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	result := make([]int, 0, n+1)
	result = append(result, 0)
	result = append(result, 1)
	result = append(result, 2)
	for i := 3; i <= n; i++ {
		result = append(result, result[i-1]+result[i-2])
	}

	return result[n]
}
