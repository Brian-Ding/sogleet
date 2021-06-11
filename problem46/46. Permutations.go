package problem46

import "fmt"

func Judge() {
	nums := []int{1, 2, 3}
	r := permute(nums)

	fmt.Println(r)
}

func permute(nums []int) [][]int {
	if len(nums) == 1 {
		return [][]int{{nums[0]}}
	}

	temp := permute(nums[:len(nums)-1])

	result := make([][]int, 0)
	for _, array := range temp {
		for i := 0; i < len(array); i++ {
			r := make([]int, 0)
			for j := 0; j < len(array); j++ {
				if j == i {
					r = append(r, nums[len(nums)-1])
				}

				r = append(r, array[j])
			}

			result = append(result, r)
		}

		result = append(result, append(array, nums[len(nums)-1]))
	}

	return result
}
