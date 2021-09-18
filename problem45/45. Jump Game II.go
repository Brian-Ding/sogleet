package problem45

import (
	"fmt"
	"math"
)

func Judge() {
	nums := []int{2, 3, 0, 1, 4}
	fmt.Println(jump(nums))
}

func jump(nums []int) int {
	// [5, 6, 4, 4, 6, 9, 4, 4, 7, 4, 4, 8, 2, 6, 8, 1, 5, 9, 6, 5, 2, 7, 9, 7, 9, 6, 9, 4, 1, 6, 8, 8, 4, 4, 2, 0, 3, 8, 5]
	//                                                                          2  2  2  2  3  2  1  1  2  2  2  f  1  1  0
	if len(nums) == 1 {
		return 0
	}

	n := len(nums)
	array := make([]int, n)
	array[n-1] = 0
	for i := n - 2; i >= 0; i-- {
		// reach the end in one jump
		if i+nums[i] >= n-1 {
			array[i] = 1
			continue
		}

		// need more jump
		min := math.MaxInt32
		for j := nums[i]; j > 0; j-- {
			jumps := array[i+j]
			if jumps != math.MaxInt32 {
				min = int(math.Min(float64(min), float64(jumps+1)))
			}
		}
		array[i] = min
	}

	return array[0]
}
