package problems

// CanJump problem 55
func CanJump(nums []int) bool {
	return canJump(nums)
}

func canJump(nums []int) bool {
	reach := make([]bool, len(nums))
	reach[0] = true
	for i := 0; i < len(nums); i++ {
		if reach[i] {
			for j := 1; j <= nums[i]; j++ {
				if i+j < len(reach) && !reach[i+j] {
					reach[i+j] = true
				}
			}
		}
	}

	return reach[len(reach)-1]
}
