package main

import (
	"fmt"

	"../problems"
)

func main() {
	a := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	r := problems.MinPathSum(a)
	fmt.Println(r)
}
