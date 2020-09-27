package main

import (
	"fmt"

	"../problems"
)

func main() {
	a := [][]int{{1, 1}, {2, 3}, {2, 3}, {1, 1}, {1, 1}, {1, 1}, {1, 2}}
	b := 4
	r := problems.MinHeightShelves(a, b)
	fmt.Println(r)
}
