package problems

// NthPersonGetsNthSeat problem 1227
func NthPersonGetsNthSeat(n int) float64 {
	return nthPersonGetsNthSeat(n)
}

func nthPersonGetsNthSeat(n int) float64 {
	if n == 1 {
		return 1
	}

	return 0.5
}
