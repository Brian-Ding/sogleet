package main

import (
	"fmt"

	"../problems"
)

func main() {
	a := "aabaaa"
	r := problems.CountSubstrings(a)
	fmt.Println(r)
}
