package problem764

import (
	"fmt"
	"math"
)

func Judge() {
	n := 5
	mines := [][]int{{3, 0}, {3, 3}}
	result := orderOfLargestPlusSign(n, mines)
	fmt.Println(result)
}

func orderOfLargestPlusSign(n int, mines [][]int) int {
	maps := make([][]int, n)
	for i := 0; i < n; i++ {
		maps[i] = make([]int, n)
		for j := 0; j < n; j++ {
			maps[i][j] = 1
		}
	}

	for i := 0; i < len(mines); i++ {
		maps[mines[i][0]][mines[i][1]] = 0
	}

	max := 0

	// left direction
	leftMaps := make([][]int, n)
	// right direction
	rightMaps := make([][]int, n)
	for i := 0; i < n; i++ {
		leftMaps[i] = make([]int, n)
		leftMaps[i][0] = maps[i][0]
		for j := 1; j < n; j++ {
			if maps[i][j] > 0 {
				leftMaps[i][j] = leftMaps[i][j-1] + 1
			}
		}

		rightMaps[i] = make([]int, n)
		rightMaps[i][n-1] = maps[i][n-1]
		for j := n - 2; j >= 0; j-- {
			if maps[i][j] > 0 {
				rightMaps[i][j] = rightMaps[i][j+1] + 1
			}
		}
	}

	// up direction
	upMaps := make([][]int, n)
	for i := 0; i < n; i++ {
		upMaps[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == 0 {
				upMaps[i][j] = maps[i][j]
			} else {
				if maps[i][j] > 0 {
					upMaps[i][j] = upMaps[i-1][j] + 1
				}
			}
		}
	}

	// down direction
	downMaps := make([][]int, n)
	for i := n - 1; i >= 0; i-- {
		downMaps[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if i == n-1 {
				downMaps[i][j] = maps[i][j]
			} else {
				if maps[i][j] > 0 {
					downMaps[i][j] = downMaps[i+1][j] + 1
				}
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			min := leftMaps[i][j]
			min = int(math.Min(float64(min), float64(rightMaps[i][j])))
			min = int(math.Min(float64(min), float64(upMaps[i][j])))
			min = int(math.Min(float64(min), float64(downMaps[i][j])))

			max = int(math.Max(float64(max), float64(min)))
		}
	}

	return max
}
