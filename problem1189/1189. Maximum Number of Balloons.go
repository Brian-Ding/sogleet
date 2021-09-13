package problem1189

import (
	"fmt"
	"math"
)

func Judge() {
	text := "nlaebolko"
	result := maxNumberOfBalloons(text)
	fmt.Println(result)
}

func maxNumberOfBalloons(text string) int {
	// balloon
	maps := make([]int, 5)
	maps[0] = 0 // b
	maps[1] = 0 // a
	maps[2] = 0 // l
	maps[3] = 0 // o
	maps[4] = 0 // n

	for i := 0; i < len(text); i++ {
		if text[i] == 'b' {
			maps[0]++
		} else if text[i] == 'a' {
			maps[1]++
		} else if text[i] == 'l' {
			maps[2]++
		} else if text[i] == 'o' {
			maps[3]++
		} else if text[i] == 'n' {
			maps[4]++
		}
	}

	result := math.MaxInt32
	result = int(math.Min(float64(result), float64(maps[0])))
	result = int(math.Min(float64(result), float64(maps[1])))
	result = int(math.Min(float64(result), float64(maps[2]/2)))
	result = int(math.Min(float64(result), float64(maps[3]/2)))
	result = int(math.Min(float64(result), float64(maps[4])))

	return result
}
