package problem72

import (
	"fmt"
	"math"
)

func Judge() {
	// word1 := "horse"
	// word2 := "ros"
	word1 := "intention"
	word2 := "execution"
	fmt.Println(minDistance(word1, word2))
}

func minDistance(word1 string, word2 string) int {
	if len(word1) == 0 && len(word2) == 0 {
		return 0
	}

	if len(word1) == 0 {
		return len(word2)
	}

	if len(word2) == 0 {
		return len(word1)
	}

	candidates := make([][]int, len(word1))
	for i := 0; i < len(word1); i++ {
		candidates[i] = make([]int, len(word2))
	}

	if word1[0] == word2[0] {
		candidates[0][0] = 0
	} else {
		candidates[0][0] = 1
	}

	for i := 1; i < len(word1); i++ {
		candidates[i][0] = candidates[i-1][0] + 1
	}

	for j := 1; j < len(word2); j++ {
		candidates[0][j] = candidates[0][j-1] + 1
	}

	for i := 1; i < len(word1); i++ {
		for j := 1; j < len(word2); j++ {
			if word1[i] == word2[j] {
				candidates[i][j] = candidates[i-1][j-1]
			} else {
				replace := candidates[i-1][j-1] + 1
				insert := candidates[i-1][j] + 1
				delete := candidates[i][j-1] + 1

				min := math.Min(float64(replace), float64(insert))
				min = math.Min(min, float64(delete))

				candidates[i][j] = int(min)
			}
		}
	}

	result := candidates[len(word1)-1][len(word2)-1]

	return result
}
