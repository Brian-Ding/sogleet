package problems

// IsSubsequence problem 392
func IsSubsequence(s string, t string) bool {
	return isSubsequence(s, t)
}

func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	}

	if t == "" {
		return false
	}

	sIndex := 0
	for _, c := range t {
		if sIndex < len(s) && byte(c) == s[sIndex] {
			sIndex++
		}
	}

	return sIndex == len(s)
}
