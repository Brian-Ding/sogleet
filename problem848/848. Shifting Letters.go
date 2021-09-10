package problem848

import "fmt"

func Judge() {
	s := "abc"
	shifts := []int{3, 5, 9}
	result := shiftingLetters(s, shifts)
	fmt.Println(result)
}

func shiftingLetters(s string, shifts []int) string {
	accumulatedShifts := make([]int, len(shifts))
	accumulatedShifts[len(shifts)-1] = shifts[len(shifts)-1]
	for i := len(shifts) - 2; i >= 0; i-- {
		accumulatedShifts[i] = shifts[i] + accumulatedShifts[i+1]
	}

	chars := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		chars[i] = shift(s[:][i], accumulatedShifts[i])
	}

	result := string(chars)

	return result
}

func shift(char byte, shift int) byte {
	// a  - b
	// 97 - 122
	delta := 123 - 97
	shift = shift % delta

	newChar := char + byte(shift)
	if newChar > 122 {
		newChar -= byte(delta)
	}

	return newChar
}
