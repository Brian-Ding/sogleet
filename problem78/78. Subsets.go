package problem78

import "fmt"

func Judge() {
	nums := []int{9, 0, 3, 5, 7}
	r := subsets(nums)

	fmt.Println(r)
}

func subsets(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums) == 1 {
		result = append(result, []int{})
		result = append(result, []int{nums[0]})

		return result
	}

	temp := subsets(nums[:len(nums)-1])
	last := nums[len(nums)-1]

	for i := 0; i < len(temp); i++ {
		new := make([]int, 0)
		for _, v := range temp[i] {
			new = append(new, v)
		}
		result = append(result, new)
		add := append(new, last)
		result = append(result, add)
	}

	return result
}
