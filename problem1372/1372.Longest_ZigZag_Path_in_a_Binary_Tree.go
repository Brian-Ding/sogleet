package problems

import "github.com/Brian-Ding/sogleet/common"

// LongestZigZag problem 1372
func LongestZigZag() int {
	a := &common.TreeNode{1, nil, &common.TreeNode{2, &common.TreeNode{3, nil, nil}, &common.TreeNode{4, &common.TreeNode{5, nil, &common.TreeNode{6, nil, &common.TreeNode{7, nil, nil}}}, &common.TreeNode{8, nil, nil}}}}
	return longestZigZag(a)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func longestZigZag(root *common.TreeNode) int {
	if root == nil {
		return 0
	}

	left := step(root.Left, true, 0)
	right := step(root.Right, false, 0)
	return common.Max(left, right)
}

func step(node *common.TreeNode, left bool, value int) int {
	if node == nil {
		return value
	}

	sum := 0
	if left {
		sum = step(node.Right, false, value+1)
		sum = common.Max(sum, step(node.Left, true, 0))
	} else {
		sum = step(node.Left, true, value+1)
		sum = common.Max(sum, step(node.Right, false, 0))
	}

	return sum
}
