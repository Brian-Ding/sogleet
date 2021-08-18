package common

import "math"

// TreeNode given by leetcode
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ListNode given by leetcode
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewArray(row, column, value int) [][]int {
	array := make([][]int, 0, row)
	for i := 0; i < row; i++ {
		array = append(array, make([]int, 0, column))
		for j := 0; j < column; j++ {
			array[i] = append(array[i], value)
		}
	}

	return array
}

func Min(nums ...int) int {
	n := math.MaxInt64
	for _, v := range nums {
		if v < n {
			n = v
		}
	}
	return n
}

func Max(nums ...int) int {
	n := math.MinInt64
	for _, v := range nums {
		if v > n {
			n = v
		}
	}
	return n
}

func Sum(nums ...int) int {
	result := 0
	for _, v := range nums {
		result += v
	}

	return result
}
