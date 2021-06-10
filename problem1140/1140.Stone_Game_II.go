package problems

// StoneGameII problem 1140
func StoneGameII(piles []int) int {
	return stoneGameII(piles)
}

// not passed
func stoneGameII(piles []int) int {
	if len(piles) <= 2 {
		return sum(piles...)
	}

	row := len(piles)
	column := row/2 + 2
	result := newArray(row, column, 0)
	for i := 0; i < row; i++ {
		for m := 1; m < column; m++ {
			if row-i <= 2*m {
				result[i][m] = sum(piles[i:]...)
			}
		}
	}

	for i := row - 1; i >= 0; i-- {
		for m := 1; m < column && row-i > 2*m; m++ {
			take := sum(piles[i:]...)
			for x := m; x <= 2*m && x < column; x++ {
				take = min(take, result[i+x][max(x, m)])
			}
			result[i][m] = sum(piles[i:]...) - take
		}
	}

	return result[0][1]
}
