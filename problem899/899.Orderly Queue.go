package problem899

import (
	"fmt"
	"sort"
	"strings"
)

func Judge() {
	s := "cba"
	k := 1
	result := orderlyQueue(s, k)
	fmt.Println(result)
}

type charSlice []rune

func (a charSlice) Len() int           { return len(a) }
func (a charSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a charSlice) Less(i, j int) bool { return a[i] < a[j] }

func orderlyQueue(s string, k int) string {
	chars := []rune(s)
	var result string
	if k == 1 {
		result = s
		for i := 1; i < len(chars); i++ {
			temp := string(s[i:]) + string(s[:i])
			if strings.Compare(result, temp) > 0 {
				result = temp
			}
		}
	} else {
		sort.Sort(charSlice(chars))
		result = string(chars)
	}

	return result
}
