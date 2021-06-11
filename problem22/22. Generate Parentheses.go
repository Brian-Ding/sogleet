package problem22

import "fmt"

func Judge() {
	r := generateParenthesis(3)

	fmt.Println(r)
}

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	result = backtrack(0, 0, n, "", result)

	return result
}

func backtrack(left, right, max int, str string, result []string) []string {
	if len(str) == 2*max {
		result = append(result, str)

		return result
	}

	if left < max {
		result = backtrack(left+1, right, max, str+"(", result)
	}

	if right < left && right < max {
		result = backtrack(left, right+1, max, str+")", result)
	}

	return result
}
