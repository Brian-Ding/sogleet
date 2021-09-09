package problem33

import "fmt"

func Judge() {
	// nums := []int{4, 5, 6, 7, 0, 1, 2}
	nums := []int{3, 1}
	target := 3
	result := search(nums, target)
	fmt.Println(result)
}

func search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	if len(nums) == 2 {
		if nums[0] == target {
			return 0
		} else if nums[1] == target {
			return 1
		} else {
			return -1
		}
	}

	start := 0
	end := len(nums)
	n := 0
	for {
		n = (start + end) / 2

		if n == start {
			break
		}

		if nums[n-1] > nums[n] {
			break
		} else if nums[0] < nums[n] {
			start = n
		} else {
			end = n
		}
	}

	if n == len(nums)-1 {
		// no rotation
		start = 0
		end = len(nums)
	} else if nums[0] < target {
		start = 0
		end = n
	} else if nums[0] > target {
		start = n
		end = len(nums)
	}

	n = binarySearch(nums[start:end], target)
	if n != -1 {
		n += start
	}

	return n
}

func binarySearch(nums []int, target int) int {
	start := 0
	end := len(nums)
	for {
		n := (start + end) / 2

		if n == start && nums[n] != target {
			return -1
		}

		if nums[n] < target {
			start = n
		} else if nums[n] == target {
			return n
		} else if nums[n] > target {
			end = n
		}
	}
}
