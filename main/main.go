package main

import (
	"fmt"

	"../problems"
)

func main() {
	a := []int{2, 7, 9, 4, 4}
	r := problems.StoneGameII(a)
	fmt.Println(r)
}
