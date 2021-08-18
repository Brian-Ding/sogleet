package problems

import (
	"math"

	"github.com/Brian-Ding/sogleet/common"
)

// MinHeightShelves problem 1105
func MinHeightShelves(books [][]int, shelf_width int) int {
	return minHeightShelves(books, shelf_width)
}

func minHeightShelves(books [][]int, shelf_width int) int {
	array := make([]int, 0, len(books))
	for i := 0; i < cap(array); i++ {
		array = append(array, math.MaxInt32)
	}

	array[0] = books[0][1]
	for i := 1; i < len(array); i++ {
		totalWdith := 0
		maxHeight := math.MinInt32
		for j := i; j >= 0; j-- {
			totalWdith += books[j][0]
			if totalWdith > shelf_width {
				break
			} else {
				maxHeight = common.Max(maxHeight, books[j][1])
				if j-1 < 0 {
					array[i] = common.Min(array[i], maxHeight)
				} else {
					array[i] = common.Min(array[i], array[j-1]+maxHeight)
				}
			}
		}
	}

	return array[len(array)-1]
}
