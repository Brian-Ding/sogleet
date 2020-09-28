package main

import (
	"fmt"

	"../problems"
)

func main() {
	a := []string{"sgtnz", "sgtz", "sgz", "ikrcyoglz", "ajelpkpx", "ajelpkpxm", "srqgtnz", "srqgotnz", "srgtnz", "ijkrcyoglz"}
	r := problems.LongestStrChain(a)
	fmt.Println(r)
}
