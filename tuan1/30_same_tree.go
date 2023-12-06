// https://leetcode.com/problems/same-tree/
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// solution 1: recursive
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val &&
		isSameTree(p.Left, q.Left) &&
		isSameTree(p.Right, q.Right)
}

// solution 2: iterative
func isSameTree(p *TreeNode, q *TreeNode) bool {
	queue := []*TreeNode{p, q}
	for len(queue) > 0 {
		p, q := queue[0], queue[1]
		queue = queue[2:]
		if p == nil || q == nil {
			if p != q {
				return false
			}
			continue
		}
		if p.Val != q.Val {
			return false
		}
		queue = append(queue, p.Left, q.Left, p.Right, q.Right)
	}
	return true
}

// solution 3: iterative
func isSameTree(p *TreeNode, q *TreeNode) bool {
	stack := []*TreeNode{p, q}
	for len(stack) > 0 {
		p, q := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]
		if p == nil || q == nil {
			if p != q {
				return false
			}
			continue
		}
		if p.Val != q.Val {
			return false
		}
		stack = append(stack, p.Left, q.Left, p.Right, q.Right)
	}
	return true
}
