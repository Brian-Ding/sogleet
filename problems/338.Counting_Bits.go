package problems

func CountBits(num int) []int {
	return countBits(num)
}

// func countBits(num int) []int {
// 	if num == 0 {
// 		return []int{0}
// 	}

// 	result := make([]int, 0, num+1)
// 	result = append(result, 0)
// 	result = append(result, 1)

// 	limit := int(math.Log2(float64(num)))
// 	for i := 1; i <= limit; i++ {
// 		temp := result[int(math.Pow(float64(2), float64(i-1))):]
// 		for _, n := range temp {
// 			result = append(result, n)
// 		}
// 		for _, n := range temp {
// 			result = append(result, n+1)
// 		}
// 	}

// 	return result[:num+1]
// }

func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}

	result := make([]int, 0, num+1)
	result = append(result, 0)

	p := 1

	for i := 1; i <= num; i++ {
		if i == 2*p {
			p = i
		}

		result = append(result, result[i-p]+1)
	}

	return result
}
