package problem199

import (
	"fmt"

	"github.com/Brian-Ding/sogleet/common"
)

func Judge() {
	root := &common.TreeNode{Val: 1, Left: &common.TreeNode{Val: 2, Right: &common.TreeNode{Val: 5}}, Right: &common.TreeNode{Val: 3, Right: &common.TreeNode{Val: 4}}}
	//root := &common.TreeNode{Val: 1, Left: &common.TreeNode{Val: 2}}
	result := rightSideView(root)

	fmt.Println(result)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *common.TreeNode) []int {
	if root == nil {
		return nil
	}

	result := bfs(root)

	return result
}

func bfs(node *common.TreeNode) []int {
	if node == nil {
		return make([]int, 0)
	}

	result := []int{node.Val}
	leftArray := make([]int, 0)
	rightArray := make([]int, 0)

	if node.Left != nil {
		leftArray = bfs(node.Left)
	}

	if node.Right != nil {
		rightArray = bfs(node.Right)
	}

	result = append(result, rightArray...)

	if len(leftArray) > len(rightArray) {
		result = append(result, leftArray[len(rightArray):]...)
	}

	return result
}
