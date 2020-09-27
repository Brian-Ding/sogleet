package main

import (
	"fmt"

	"../problems"
)

func main() {
	a := []int{1, 2, 3, 4, 6, 8, 9, 10, 13, 14, 16, 17, 19, 21, 24, 26, 27, 28, 29}
	b := []int{3, 14, 50}
	r := problems.MincostTickets(a, b)
	fmt.Println(r)
}
