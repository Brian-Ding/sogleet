package problems

import "sort"

// FindLongestChain problem 646
func FindLongestChain() int {
	a := [][]int{{1, 2}, {2, 3}, {3, 4}}
	return findLongestChain(a)
}

func findLongestChain(pairs [][]int) int {
	sort.Sort(byNum(pairs))
	width := len(pairs)
	array := make([]int, 0, len(pairs))
	for i := 0; i < width; i++ {
		array = append(array, 1)
	}

	for i := 1; i < width; i++ {
		for j := 0; j < i; j++ {
			if pairs[i][0] > pairs[j][1] {
				array[i] = max(array[i], array[j]+1)
			}
		}
	}

	return array[width-1]
}

type byNum [][]int

func (s byNum) Len() int {
	return len(s)
}
func (s byNum) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s byNum) Less(i, j int) bool {
	return s[i][0] < s[j][0]
}
