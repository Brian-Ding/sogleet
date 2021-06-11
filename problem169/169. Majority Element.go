package problem169

import (
	"fmt"
)

func Judge() {
	// nums := []int{3,2,3}
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	r := majorityElement(nums)

	fmt.Println(r)
}

// Given an array nums of size n, return the majority element.
// The majority element is the element that appears more than ⌊n / 2⌋ times.
// You may assume that the majority element always exists in the array.

// Follow-up: Could you solve the problem in linear time and in O(1) space?

func majorityElement(nums []int) int {
	dict := make(map[int]int, 0)

	for i := 0; i < len(nums); i++ {
		if val, ok := dict[nums[i]]; ok {
			dict[nums[i]] = val + 1
		} else {
			dict[nums[i]] = 1
		}
	}

	max := 0
	key := 0
	for k, v := range dict {
		if v > max {
			key = k
			max = v
		}
	}

	return key
}
