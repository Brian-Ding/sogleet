package main

import (
	"fmt"

	"../problems"
)

func main() {
	a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	r := problems.MinFallingPathSum(a)
	fmt.Println(r)
}
