package main

import (
	"fmt"
	"problems"
)

func main() {
	a := []int{1, 15, 7, 9, 2, 5, 10}
	r := problems.MaxSumAfterPartitioning(a, 3)
	fmt.Println(r)
}
