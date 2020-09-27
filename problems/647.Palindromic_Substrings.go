package problems

// CountSubstrings problem 647
func CountSubstrings(s string) int {
	return countSubstrings(s)
}

func countSubstrings(s string) int {
	array := make([]byte, 0, len(s)*2-1)
	for i := 0; i < cap(array); i++ {
		if i%2 == 0 {
			array = append(array, s[i/2])
		} else {
			array = append(array, '*')
		}
	}
	sum := len(s)
	for i := 1; i < len(array)-1; i++ {
		start := 0
		if array[i] == '*' {
			start = 1
		} else {
			start = 2
		}
		for j := start; i-j >= 0 && i+j < len(array); j += 2 {
			if array[i-j] == array[i+j] {
				sum++
			} else {
				break
			}
		}
	}

	return sum
}
