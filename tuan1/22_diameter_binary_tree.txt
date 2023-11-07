// https://leetcode.com/problems/diameter-of-binary-tree/
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2, Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	println(diameterOfBinaryTree(root))
}

func diameterOfBinaryTree(root *TreeNode) int {
	var max = 0
	depth(root, &max)
	return max
}

func depth(root *TreeNode, max *int) int {
	if root == nil {
		return 0
	}
	left := depth(root.Left, max)
	right := depth(root.Right, max)
	if left+right > *max {
		*max = left + right
	}
	if left > right {
		return left + 1
	}
	return right + 1
}
