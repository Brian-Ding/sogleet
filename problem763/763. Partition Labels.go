package problems

import "strings"

func PartitionLabels() []int {
	s := "ababcbacadefegdehijhklij"
	// a:0	8
	// b:1	5
	// c:4	7
	// d:9	14
	// e:10	15
	// f:11	11
	// g:13	13
	// h:16	19
	// i:17	22
	// j:18	23
	// k:20	20
	// l:21	21
	return partitionLabels(s)
}

func partitionLabels(s string) []int {
	if len(s) == 0 {
		return make([]int, 0)
	}
	byteStr := []byte(s)
	// maps := make(map[byte][]int, 0)
	// for i := 0; i < len(s); i++ {
	// 	v, ok := maps[byteStr[i]]
	// 	if ok {
	// 		maps[byteStr[i]] = []int{v[0], i}
	// 	} else {
	// 		maps[byteStr[i]] = []int{i, i}
	// 	}
	// }

	// result := make([]int, 1)

	// for i := 1; i < len(byteStr); i++ {
	// 	c := byteStr[i]
	// 	v0 := &maps[c][0]
	// 	v1 := &maps[c][1]
	// 	if v[0] <= max && v[1] > max {
	// 		max = v[1]
	// 	} else if v[0] > max && v[1] > max {
	// 		result = append(result, max)
	// 		max = v[1]
	// 	}
	// }

	// return result

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
