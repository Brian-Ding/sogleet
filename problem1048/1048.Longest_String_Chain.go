package problems

import "sort"

// LongestStrChain problem 1048
func LongestStrChain(words []string) int {
	return longestStrChain(words)
}

func longestStrChain(words []string) int {
	sort.Sort(ByLength(words))
	array := make([]int, 0, len(words))
	for i := 0; i < cap(array); i++ {
		array = append(array, 1)
	}
	result := make([]int, 0, len(words))
	for i := 0; i < cap(result); i++ {
		result = append(result, 1)
	}

	for i := 1; i < len(array); i++ {
		result[i] = result[i-1]
		for j := i - 1; j >= 0; j-- {
			if len(words[i])-len(words[j]) == 1 {
				diff := 0
				for k := 0; k < len(words[j]); k++ {
					if diff > 1 {
						break
					}
					if words[i][k+diff] != words[j][k] {
						diff++
						k--
					}
				}

				if diff <= 1 {
					array[i] = max(array[i], array[j]+1)
					result[i] = max(result[i], array[i])
				}
			}
		}
	}

	return result[len(result)-1]
}

// ByLength sort string slice
type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}
