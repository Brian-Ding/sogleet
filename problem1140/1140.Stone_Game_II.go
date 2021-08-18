package problems

import "github.com/Brian-Ding/sogleet/common"

// StoneGameII problem 1140
func StoneGameII(piles []int) int {
	return stoneGameII(piles)
}

// not passed
func stoneGameII(piles []int) int {
	if len(piles) <= 2 {
		return common.Sum(piles...)
	}

	row := len(piles)
	column := row/2 + 2
	result := common.NewArray(row, column, 0)
	for i := 0; i < row; i++ {
		for m := 1; m < column; m++ {
			if row-i <= 2*m {
				result[i][m] = common.Sum(piles[i:]...)
			}
		}
	}

	for i := row - 1; i >= 0; i-- {
		for m := 1; m < column && row-i > 2*m; m++ {
			take := common.Sum(piles[i:]...)
			for x := m; x <= 2*m && x < column; x++ {
				take = common.Min(take, result[i+x][common.Max(x, m)])
			}
			result[i][m] = common.Sum(piles[i:]...) - take
		}
	}

	return result[0][1]
}
