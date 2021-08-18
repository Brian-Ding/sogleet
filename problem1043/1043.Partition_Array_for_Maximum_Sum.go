package problems

import (
	"math"

	"github.com/Brian-Ding/sogleet/common"
)

// MaxSumAfterPartitioning problem 1043
func MaxSumAfterPartitioning(arr []int, k int) int {
	return maxSumAfterPartitioning(arr, k)
}

func maxSumAfterPartitioning(arr []int, k int) int {
	result := make([]int, 0, len(arr))
	result = append(result, arr[0])
	result = append(result, common.Max(arr[0], arr[1])*2)
	result = append(result, common.Max(arr[0], arr[1], arr[2])*3)

	for i := 3; i < len(arr); i++ {
		temp := math.MinInt32
		for j := 1; j <= 3; j++ {
			temp = common.Max(result[i-j]+j*common.Max(arr[i-j+1:i+1]...), temp)
		}
		result = append(result, temp)
	}

	return result[len(arr)-1]
}
