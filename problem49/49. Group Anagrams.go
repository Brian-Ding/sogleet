package problem49

import "fmt"

func Judge() {
	// Example 1:
	// Input: strs = ["eat","tea","tan","ate","nat","bat"]
	// Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

	// Example 2:

	// Input: strs = [""]
	// Output: [[""]]

	// Example 3:
	// Input: strs = ["a"]
	// Output: [["a"]]

	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(strs)
	fmt.Println(result)
}

func groupAnagrams(strs []string) [][]string {
	dict := make(map[string][]string)

	for _, str := range strs {
		found := false
		for key, value := range dict {
			if compare(key, str) {
				value = append(value, str)
				found = true
			}
		}
		if !found {
			dict[str] = make([]string, 0)
			dict[str] = append(dict[str], str)
		}
	}

	result := make([][]string, 0)
	for _, value := range dict {
		result = append(result, value)
	}

	return result
}

func compare(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	// eat tea
	return false
}
