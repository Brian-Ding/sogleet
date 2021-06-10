package problems

import "math"

func LastStoneWeightII() int {
	stones := []int{2, 7, 4, 1, 8, 1}
	return lastStoneWeightII(stones)
}

func lastStoneWeightII(stones []int) int {
	// all possible from sub stones
	// match the sub stones combination with closest value

	// out of memory
	if len(stones) == 0 {
		return 0
	}

	if len(stones) == 1 {
		return stones[0]
	}

	matches := make([][]int, 0, len(stones))
	matches = append(matches, []int{stones[0]})
	for i := 1; i < len(stones); i++ {
		values := make([]int, 0)
		for j := 0; j < len(matches[i-1]); j++ {
			abs := int(math.Abs(float64(matches[i-1][j] + stones[i])))
			has := false
			for _, v := range values {
				if v == abs {
					has = true
					break
				}
			}
			if !has {
				values = append(values, abs)
			}

			abs = int(math.Abs(float64(matches[i-1][j] - stones[i])))
			has = false
			for _, v := range values {
				if v == abs {
					has = true
					break
				}
			}
			if !has {
				values = append(values, abs)
			}
		}

		matches = append(matches, values)
	}

	minimum := math.MaxInt32
	for _, v := range matches[len(matches)-1] {
		minimum = int(math.Min(math.Abs(float64(v)), math.Abs(float64(minimum))))
	}

	return minimum
}
