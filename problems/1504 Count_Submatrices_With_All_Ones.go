package problems

/*
Given a rows * columns matrix mat of ones and zeros, return how many submatrices have all ones.
*/

func NumSubmat(mat [][]int) int {
	return numSubmat(mat)
}

// a := [][]int{{0, 1, 1, 0}, {0, 1, 1, 1}, {1, 1, 1, 0}}
// a := [][]int{{1, 0, 1, 1}, {1, 1, 1, 0}, {1, 1, 1, 1}, {1, 1, 1, 1}}
// a := [][]int{{1, 0, 0, 1, 0, 1, 0, 1}, {1, 0, 1, 1, 0, 1, 0, 0}, {1, 1, 1, 0, 1, 0, 0, 1}, {0, 0, 1, 1, 1, 1, 0, 0}, {0, 0, 1, 1, 1, 1, 0, 1}, {1, 1, 0, 1, 1, 1, 0, 0}}
// not pass
func numSubmat(mat [][]int) int {
	row := len(mat)
	column := len(mat[0])

	array := make([][]point, 0, row)
	for i := 0; i < row; i++ {
		array = append(array, make([]point, 0, column))
		for j := 0; j < column; j++ {
			array[i] = append(array[i], point{})
		}
	}

	for i := 1; i < row; i++ {
		if mat[i][0] != 0 && mat[i-1][0] != 0 {
			array[i][0].top = array[i-1][0].top + 1
		}
	}

	for j := 1; j < column; j++ {
		if mat[0][j] != 0 && mat[0][j-1] != 0 {
			array[0][j].left = array[0][j-1].left + 1
		}
	}

	for i := 1; i < row; i++ {
		for j := 1; j < column; j++ {
			if mat[i][j] != 0 {
				if mat[i-1][j] != 0 {
					array[i][j].top = array[i-1][j].top + 1
				}
				if mat[i][j-1] != 0 {
					array[i][j].left = array[i][j-1].left + 1
				}
				if mat[i-1][j] != 0 && mat[i][j-1] != 0 && mat[i-1][j-1] != 0 {
					h := array[i][j].top
					l := array[i][j].left
					h2 := h
					l2 := l
					for k := 1; i-k >= 0 && j-k >= 0; k++ {
						p := array[i-k][j-k]
						h = min(p.top+k, h)
						l = min(p.left+k, l)

						if mat[i-k][j-k] == 0 {
							h2 = k - 1
							l2 = k - 1
							break
						}
					}

					array[i][j].leftTop = h*l2 + h2*l - h2*l2
				}
			}
		}
	}

	sum := 0
	for i := 0; i < row; i++ {
		for j := 0; j < column; j++ {
			sum += mat[i][j] + array[i][j].left + array[i][j].top + array[i][j].leftTop
		}
	}

	return sum
}

type point struct {
	left    int
	top     int
	leftTop int
}
