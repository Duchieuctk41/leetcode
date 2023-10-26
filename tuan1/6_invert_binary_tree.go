// https://leetcode.com/problems/invert-binary-tree/
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

func invertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	// swap left and right
	root.Left, root.Right = root.Right, root.Left

	// invert left and right
	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

func invertTree2(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	}
	return root
}
