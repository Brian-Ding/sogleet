package problems

import "strings"

func PartitionLabels() []int {
	s := "ababcbacadefegdehijhklij"
	return partitionLabels(s)
}

func partitionLabels(s string) []int {
	if len(s) == 0 {
		return make([]int, 0)
	}
	byteStr := []byte(s)
	results := make(map[int][]int, 0)
	results[0] = []int{0}

	for i := 1; i < len(s); i++ {
		results[i] = make([]int, 0)
		first := strings.IndexByte(s, byteStr[i])
		if first < i {
			// this char shows at least once before
			//ababcbaca defegde hijh k l
			//ababcbaca defegde hijhkli
			//ababcbaca defegde hijhklij
			for _, v := range results[i-1] {
				if first > v {
					results[i] = append(results[i], v)
				} else {
					// character byteStr[i] must be grouped with v
					results[i] = append(results[i], i)
					break
				}
			}
		} else {
			results[i] = results[i-1]
			results[i] = append(results[i], i)
		}
	}

	lastIndex := len(s) - 1
	result := make([]int, 0, len(results[lastIndex]))
	result = append(result, results[lastIndex][0]+1)
	for i := 1; i < cap(result); i++ {
		result = append(result, results[lastIndex][i]-results[lastIndex][i-1])
	}

	return result
}
