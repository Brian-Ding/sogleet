package problems

import "math"

func MaxProduct(ns []int) int {
	if len(ns) == 0 {
		return 0
	}

	minv, maxv, result := ns[0], ns[0], ns[0]

	for _, v := range ns[1:] {
		minv, maxv = min(v, maxv*v, minv*v), max(v, maxv*v, minv*v)
		result = max(maxv, result)
	}

	return result
}

func min(vs ...int) int {
	n := math.MaxInt64
	for _, v := range vs {
		if v < n {
			n = v
		}
	}
	return n
}

func max(vs ...int) int {
	n := math.MinInt64
	for _, v := range vs {
		if v > n {
			n = v
		}
	}
	return n
}
