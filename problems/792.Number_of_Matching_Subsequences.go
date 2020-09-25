package problems

func NumMatchingSubseq(S string, words []string) int {
	return numMatchingSubseq(S, words)
}

func numMatchingSubseq(S string, words []string) int {
	results := make(map[string]int)
	for _, s := range words {
		results[s] = 0
	}

	for _, c := range S {
		for k, v := range results {
			if v < len(k) && byte(c) == k[v] {
				results[k]++
			}
		}
	}

	count := 0
	for k, v := range results {
		if v == len(k) {
			count++
		}
	}

	return count
}
