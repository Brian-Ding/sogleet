package main

import (
	"fmt"

	"../problems"
)

func main() {
	a := []int{1, 2, 3, 4}
	r := problems.NumberOfArithmeticSlices(a)
	fmt.Println(r)
}
