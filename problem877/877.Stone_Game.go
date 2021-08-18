package problems

import (
	"math"

	"github.com/Brian-Ding/sogleet/common"
)

// StoneGame problem 877
func StoneGame(piles []int) bool {
	return stoneGame(piles)
}

func stoneGame(piles []int) bool {
	length := len(piles)
	firstArray := common.NewArray(length, length, 0)
	secondArray := common.NewArray(length, length, 0)

	for i := 0; i < length; i++ {
		firstArray[i][i] = piles[i]
	}

	for delta := 1; delta < length; delta++ {
		for i := 0; i+delta < length; i++ {
			left := piles[i] + secondArray[i+1][i+delta]
			right := piles[i+delta] + secondArray[i][i+delta-1]
			firstArray[i][i+delta] = int(math.Max(float64(left), float64(right)))
		}
	}

	return firstArray[0][length-1] > secondArray[0][length-1]
}
