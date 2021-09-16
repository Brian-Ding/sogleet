package problem978

import "fmt"

func Judge() {
	arr := []int{4, 8, 12, 16}
	fmt.Println(maxTurbulenceSize(arr))
}

func maxTurbulenceSize(arr []int) int {
	// 9, 4, 2, 10, 7, 8, 8, 1, 9
	if len(arr) == 1 {
		return 1
	}

	if len(arr) == 2 {
		if arr[0] == arr[1] {
			return 1
		} else {
			return 2
		}
	}

	// candidate maximum length
	candidates := make([]int, len(arr))
	candidates[0] = 1
	if arr[0] == arr[1] {
		candidates[1] = 1
	} else {
		candidates[1] = 2
	}

	result := 1
	for i := 2; i < len(arr); i++ {
		if arr[i-1] == arr[i] {
			candidates[i] = 1
			continue
		}

		if (arr[i]-arr[i-1])*(arr[i-1]-arr[i-2]) < 0 {
			candidates[i] = candidates[i-1] + 1
		} else {
			candidates[i] = 2
		}

		if candidates[i] > result {
			result = candidates[i]
		}
	}

	return result
}
