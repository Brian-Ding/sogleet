package problems

import (
	"math"
)

// Rob2 problem 213
func Rob2(nums []int) int {
	return rob2(nums)
}

func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	nums1 := nums[:len(nums)-1]
	nums2 := nums[1:]

	robResult1 := make([]int, len(nums)-1)
	robResult2 := make([]int, len(nums)-1)

	for i := 0; i < len(nums1); i++ {
		if i-2 < 0 && i-1 < 0 { // i = 0
			robResult1[i] = nums1[0]
			robResult2[i] = nums2[0]
		} else if i-2 < 0 { // i = 1
			robResult1[i] = int(math.Max(float64(nums1[0]), float64(nums1[1])))
			robResult2[i] = int(math.Max(float64(nums2[0]), float64(nums2[1])))
		} else { // i >= 2
			robResult1[i] = int(math.Max(float64(robResult1[i-2]+nums1[i]), float64(robResult1[i-1])))
			robResult2[i] = int(math.Max(float64(robResult2[i-2]+nums2[i]), float64(robResult2[i-1])))
		}
	}

	return int(math.Max(float64(robResult1[len(robResult1)-1]), float64(robResult2[len(robResult2)-1])))
}
