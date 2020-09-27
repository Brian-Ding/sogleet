package main

import (
	"fmt"

	"../problems"
)

func main() {
	a := "delete"
	b := "leet"
	r := problems.LongestCommonSubsequence(a, b)
	fmt.Println(r)
}
