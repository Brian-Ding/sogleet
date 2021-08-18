package problems

import "github.com/Brian-Ding/sogleet/common"

// MaxProduct problem 152
func MaxProduct(ns []int) int {
	return maxProduct(ns)
}

func maxProduct(ns []int) int {
	if len(ns) == 0 {
		return 0
	}

	minv, maxv, result := ns[0], ns[0], ns[0]

	for _, v := range ns[1:] {
		minv, maxv = common.Min(v, maxv*v, minv*v), common.Max(v, maxv*v, minv*v)
		result = common.Max(maxv, result)
	}

	return result
}
