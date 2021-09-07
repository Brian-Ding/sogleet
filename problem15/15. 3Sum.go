package problem15

import (
	"fmt"
	"sort"
)

func Judge() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum(nums)
	fmt.Println(result)
}

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	sort.Ints(nums)

	result := make([][]int, 0)
	n := len(nums)
	for i := 0; i < n-2; i++ {
		for i != 0 && i < n-2 && nums[i] == nums[i-1] {
			i++
		}

		if i >= n-2 {
			break
		}

		j := i + 1
		k := n - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum < 0 {
				j++
			} else if sum == 0 {
				result = append(result, []int{nums[i], nums[j], nums[k]})
				b := nums[j]
				c := nums[k]
				for j < k && b == nums[j] {
					j++
				}
				for j < k && c == nums[k] {
					k--
				}
			} else {
				k--
			}
		}
	}

	return result
}
