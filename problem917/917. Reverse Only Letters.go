package problem917

import (
	"fmt"
)

func Judge() {
	// s := "ab-cd"
	// s := "a-bC-dEf-ghIj"
	// s := "Test1ng-Leet=code-Q!"
	s := "7_28]"
	result := reverseOnlyLetters(s)
	fmt.Println(result)
}

func reverseOnlyLetters(s string) string {
	bytes := []byte(s)
	start := 0
	end := len(bytes) - 1
	for {
		if start >= end {
			break
		}

		if !isLetter(bytes[start]) {
			start++
		} else if !isLetter(bytes[end]) {
			end--
		} else {
			bytes[start], bytes[end] = bytes[end], bytes[start]
			start++
			end--
		}
	}

	result := string(bytes)

	return result
}

func isLetter(b byte) bool {
	result := (65 <= b && b <= 90) || (97 <= b && b <= 122)

	return result
}
