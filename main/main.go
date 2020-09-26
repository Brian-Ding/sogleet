package main

import (
	"fmt"

	"../problems"
)

func main() {
	a := []int{1, 4, 6, 7, 8, 20}
	b := []int{2, 7, 15}
	r := problems.MincostTickets(a, b)
	fmt.Println(r)
}
