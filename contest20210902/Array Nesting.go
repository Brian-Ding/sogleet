package contest20210902

import "fmt"

func Judge() {
	nums := []int{5, 4, 0, 3, 1, 6, 2}
	result := arrayNesting(nums)
	fmt.Println(result)
}

func arrayNesting(nums []int) int {
	//0, 1, 2, 3, 4, 5, 6
	//5, 4, 0, 3, 1, 6, 2			//5, 6, 2, 0

	result := 0

	for i := 0; i < len(nums); i++ {
		indicators := make([]bool, len(nums))
		candidate := 0

		k := i
		for {
			if !indicators[k] {
				indicators[k] = true
				candidate++
				k = nums[k]
			} else {
				break
			}
		}

		if candidate > result {
			result = candidate
		}
	}

	return result
}
