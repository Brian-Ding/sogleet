package problem1049

import (
	"fmt"
	"math"
)

func Judge() {
	//stones := []int{2, 7, 4, 1, 8, 1}
	stones := []int{1, 2}
	r := lastStoneWeightII(stones)
	fmt.Println(r)
}

func lastStoneWeightII(stones []int) int {
	//r1 := solveWithBinaryTree(stones)
	r := solveWithBag(stones)
	return r
}

func solveWithBinaryTree(stones []int) int {
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

func solveWithBag(stones []int) int {
	if len(stones) == 0 {
		return 0
	}

	if len(stones) == 1 {
		return stones[0]
	}

	sum := 0
	for _, v := range stones {
		sum += v
	}

	full := int(math.Floor(float64(sum) / 2))
	// len(stones) * (full + 1)
	bag := make([][]int, len(stones))
	for i := 0; i < len(bag); i++ {
		bag[i] = make([]int, full+1)
	}

	for j := 1; j < full+1; j++ {
		if stones[0] <= j {
			bag[0][j] = stones[0]
		}
	}

	for i := 1; i < len(stones); i++ {
		for j := 1; j < full+1; j++ {
			if stones[i] <= j {
				next := bag[i-1][j-stones[i]] + stones[i]
				bag[i][j] = int(math.Max(float64(next), float64(bag[i-1][j])))
			} else {
				bag[i][j] = bag[i-1][j]
			}
		}
	}

	result := sum - 2*bag[len(stones)-1][full]
	return result
}
