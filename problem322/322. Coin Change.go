package problem322

import (
	"fmt"
	"math"
)

func Judge() {
	coins := []int{2}
	amount := 3
	fmt.Println(coinChange(coins, amount))
}

func coinChange(coins []int, amount int) int {
	result := change(coins, amount)

	if result == math.MaxInt32 {
		result = -1
	}

	return result
}

func change(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	min := math.MaxInt32
	for i := 0; i < len(coins); i++ {
		left := amount - coins[i]
		candidate := math.MaxInt32
		if left > 0 {
			result := coinChange(coins, left)
			if result == -1 {
				candidate = -1
			} else {
				candidate = result + 1
			}
		} else if left == 0 {
			candidate = 1
		}

		if candidate < min {
			min = candidate
		}
	}

	return min
}

// func coinChange(coins []int, amount int) int {
// 	if amount == 0 {
// 		return 0
// 	}

// 	array := make([]int, amount+1)
// 	array[0] = -1
// 	for i := 0; i < len(coins); i++ {
// 		index := coins[i]
// 		if index < len(array) {
// 			array[coins[i]] = 1
// 		}
// 	}

// 	for i := 1; i < len(array); i++ {
// 		if array[i] == 1 {
// 			continue
// 		}

// 		min := math.MaxInt
// 		for n := 0; n < len(coins); n++ {
// 			index := i - coins[n]
// 			if index > 0 && array[index] != -1 {
// 				min = int(math.Min(float64(min), float64(array[index]+1)))
// 			}
// 		}

// 		if min == math.MaxInt {
// 			array[i] = -1
// 		} else {
// 			array[i] = min
// 		}

// 	}

// 	return array[amount]
// }
