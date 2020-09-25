package problems

import (
	"math"
)

func MaxSumAfterPartitioning(arr []int, k int) int {
	return maxSumAfterPartitioning(arr, k)
}

func maxSumAfterPartitioning(arr []int, k int) int {
	result := make([]int, 0, len(arr))
	result = append(result, arr[0])
	result = append(result, getMax(arr[0], arr[1])*2)
	result = append(result, getMax(arr[0], arr[1], arr[2])*3)

	for i := 3; i < len(arr); i++ {
		max := math.MinInt32
		for j := 1; j <= 3; j++ {
			temp := result[i-j] + j*getMax(arr[i-j+1:i+1]...)
			if temp > max {
				max = temp
			}
		}
		result = append(result, max)
	}

	return result[len(arr)-1]
}

func getMax(nums ...int) int {
	max := math.MinInt32
	for _, n := range nums {
		if n > max {
			max = n
		}
	}

	return max
}
