package contest20210816

import (
	"fmt"
	"strings"
)

func Judge() {
	//"A"				"A"								""
	//"AD"				"A"								""
	//"ADO"				"A"								""
	//"ADOB"			"AB"							""
	//"ADOBE"			"AB"							""
	//"ADOBEC"			"ABC"							"ADOBEC"
	//"ADOBECO"			"ABC"							"ADOBEC"
	//"ADOBECOD"		"ABC"							"ADOBEC"
	//"ADOBECODE"		"ABC"							"ADOBEC"
	//"ADOBECODEB"		"ABCB"							"ADOBEC"
	//"ADOBECODEBA"		"ABCBA"							"ADOBEC","CODEBA"
	//"ADOBECODEBAN"	"ABCBA"							"ADOBEC","CODEBA"
	//"ADOBECODEBANC"	"ABCBAC"						"BANC"

	// A: 0, 10
	// B: 3, 9
	// C: 5, 12

	// s := "ADOBECODEBANC"
	// t := "ABC"
	s := "a"
	t := "b"

	result := minWindow(s, t)
	fmt.Println(result)
}

func minWindow(s string, t string) string {
	if s == "" || t == "" {
		return ""
	}

	if len(s) < len(t) {
		return ""
	}

	maps := make(map[rune]int)
	for _, v := range t {
		if _, ok := maps[v]; ok {
			maps[v]++
		} else {
			maps[v] = 1
		}
	}

	left := 0
	right := len(t) - 1
	minLeft := -1
	minRight := len(s) - 1

	for {
		if right-left+1 >= len(t) && right < len(s) {
			if overlap(s[left:right+1], t) {
				if right-left < minRight-minLeft {
					minLeft = left
					minRight = right
				}

				left++
			} else {
				right++
			}
		} else {
			break
		}
	}

	if minLeft == -1 {
		return ""
	}

	result := s[minLeft : minRight+1]

	return result
}

func overlap(s string, t string) bool {
	// temp := t
	// for _, v := range s {
	// 	index := strings.Index(temp, string(v))
	// 	if index >= 0 {
	// 		temp = temp[:index] + temp[index+1:]
	// 	}
	// }

	// return temp == ""

	for _, v := range t {
		newS := strings.ReplaceAll(s, string(v), "")
		newT := strings.ReplaceAll(t, string(v), "")

		if len(s)-len(newS) < len(t)-len(newT) {
			return false
		}
	}

	return true
}
