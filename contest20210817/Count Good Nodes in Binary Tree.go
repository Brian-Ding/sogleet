package contest20210817

import (
	"fmt"

	"github.com/Brian-Ding/sogleet/common"
)

func Judge() {
	root := &common.TreeNode{Val: 3, Left: &common.TreeNode{Val: 1, Left: &common.TreeNode{Val: 3}}, Right: &common.TreeNode{Val: 4, Left: &common.TreeNode{Val: 1}, Right: &common.TreeNode{Val: 5}}}
	result := goodNodes(root)

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
func goodNodes(root *common.TreeNode) int {
	result := iterate(root, root.Val)

	return result
}

func iterate(node *common.TreeNode, max int) int {
	if node == nil {
		return 0
	}

	total := 0
	if node.Val >= max {
		max = node.Val
		total++
	}

	total += iterate(node.Left, max)
	total += iterate(node.Right, max)

	return total
}
