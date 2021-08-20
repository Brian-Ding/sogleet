package problem91

import "fmt"

func Judge() {
	s := "1201234"
	result := numDecodings(s)

	fmt.Println(result)
}

func numDecodings(s string) int {
	if len(s) == 0 || s == "0" || s[0] == '0' {
		return 0
	}

	maps := make(map[int][][]int)
	maps[0] = [][]int{{int(s[0]) - 48}}
	for i := 1; i < len(s); i++ {
		maps[i] = make([][]int, 0)
		if int(s[i])-48 != 0 {
			for _, v := range maps[i-1] {
				temp := v[:]
				temp = append(temp, int(s[i])-48)
				maps[i] = append(maps[i], temp)
			}
		}

		for _, v := range maps[i-1] {
			if v[len(v)-1]*10+int(s[i])-48 <= 26 && int(s[i-1])-48 != 0 {
				temp := v[:len(v)-1]
				temp = append(temp, v[len(v)-1]*10+int(s[i])-48)
				maps[i] = append(maps[i], temp)
			}
		}
	}

	result := len(maps[len(s)-1])

	return result
}
