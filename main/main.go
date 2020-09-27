package main

import (
	"fmt"

	"../problems"
)

func main() {
	a := "delete"
	b := "leet"
	r := problems.MinimumDeleteSum(a, b)
	fmt.Println(r)
}
