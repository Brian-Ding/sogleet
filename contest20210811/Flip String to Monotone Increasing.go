package contest20210811

import (
	"fmt"
	"math"
)

func Judge() {
	//s := "0" // expect 0
	//s := "00" // expect 0
	//s := "001" // expect 0
	//s := "0011" // expect 0
	//s := "00110" // expect 1

	//s := "0" // expect 0
	//s := "01" // expect 0
	//s := "010" // expect 1
	//s := "0101" // expect 1
	//s := "01011" // expect 1
	//s := "010110" // expect 2

	//s := "0" // expect 0
	//s := "00" // expect 0
	//s := "000" // expect 0
	//s := "0001" // expect 0
	//s := "00011" // expect 0
	//s := "000110" // expect 1
	//s := "0001100" // expect 2
	//s := "00011000" // expect 2

	//s := "1" // expect 0
	//s := "10" // expect 1
	//s := "100" // expect 1
	//s := "1001" // expect 1
	//s := "10011" // expect 1
	//s := "100111" // expect 1
	//s := "1001111" // expect 1
	//s := "10011111" // expect 1
	//s := "100111111" // expect 1
	//s := "1001111111" // expect 1
	//s := "10011111110" // expect 2
	//s := "100111111100" // expect 3
	//s := "1001111111001" // expect 3
	//s := "10011111110010" // expect 4
	//s := "100111111100101" // expect 4
	//s := "1001111111001011" // expect 4
	//s := "10011111110010111" // expect 4
	//s := "100111111100101110" // expect 5
	//s := "1001111111001011101" // expect 5
	//s := "10011111110010111011" // expect 5
	//s := "100111111100101110" // expect 5

	//s := "0" // expect 0
	//s := "01" // expect 0
	//s := "011" // expect 0
	//s := "0111" // expect 0
	s := "01110" // expect 1
	result := minFlipsMonoIncr(s)

	fmt.Println(result)
}

func minFlipsMonoIncr(s string) int {
	if len(s) == 0 || len(s) == 1 {
		return 0
	}

	if len(s) == 2 && s[0] != s[1] {
		return 1
	}

	ones := make([]int, len(s))
	if s[0] == '1' {
		ones[0] = 1
	}

	array := make([]int, len(s))
	for i := 1; i < len(s); i++ {
		ones[i] = ones[i-1]
		if s[i] == '1' {
			ones[i]++
			array[i] = array[i-1]
		} else {
			a := array[i-1] + 1 // flip this digit
			b := ones[i-1]      // flip all ones before it
			array[i] = int(math.Min(float64(a), float64(b)))
		}
	}

	result := array[len(s)-1]

	return result
}
