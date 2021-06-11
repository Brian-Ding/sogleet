package problem94

import (
	"fmt"

	common "github.com/Brian-Ding/sogleet/common"
)

func Judge() {
	root := &common.TreeNode{}
	r := inorderTraversal(root)

	fmt.Println(r)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *common.TreeNode) []int {
	result := make([]int, 0)
	result = scan(root, result)
	return result
}

func scan(root *common.TreeNode, input []int) []int {
	if root == nil {
		return input
	}

	result := scan(root.Left, input)
	result = append(result, root.Val)
	result = scan(root.Right, result)

	return result
}
