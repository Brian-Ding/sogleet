package problems

func Fib(N int) int {
	return fib(N)
}

func fib(N int) int {
	if N < 2 {
		return N
	}

	a := 0
	b := 1

	for i := 2; i <= N; i++ {
		a, b = b, a+b
	}

	return b
}
