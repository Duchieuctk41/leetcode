// https://leetcode.com/problems/subtree-of-another-tree/
package main

import (
	"strconv"
	"strings"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// Solution 1: using string
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	s1 := preorder(root)
	s2 := preorder(subRoot)
	return strings.Contains(s1, s2)
}

func preorder(root *TreeNode) string {
	if root == nil {
		return "null"
	}
	return "#" + strconv.Itoa(root.Val) + preorder(root.Left) + preorder(root.Right)
}

// Solution 2: using recursion
func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil {
		return false
	}
	if isSameTree(root, subRoot) {
		return true
	}
	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}
	if root == nil || subRoot == nil {
		return false
	}
	if root.Val != subRoot.Val {
		return false
	}
	return isSameTree(root.Left, subRoot.Left) && isSameTree(root.Right, subRoot.Right)
}
