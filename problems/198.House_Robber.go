package problems

import (
	"math"
)

func Rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	robResult := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		if i-2 < 0 && i-1 < 0 { // i = 0
			robResult[i] = nums[0]
		} else if i-2 < 0 { // i = 1
			robResult[i] = int(math.Max(float64(nums[0]), float64(nums[1])))
		} else { // i >= 2
			robResult[i] = int(math.Max(float64(robResult[i-2]+nums[i]), float64(robResult[i-1])))
		}
	}

	return robResult[len(nums)-1]
}
