// https://leetcode.com/problems/symmetric-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// solution 1: recursive
func isSymmetric(root *TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(t1 *TreeNode, t2 *TreeNode) bool {
	// 1. Both empty
	if t1 == nil && t2 == nil {
		return true
	}
	// 2. Only one empty
	if t1 == nil || t2 == nil {
		return false
	}
	// 3. Both non-empty
	return t1.Val == t2.Val && isMirror(t1.Left, t2.Right) && isMirror(t1.Right, t2.Left)
}

// solution 2: iterative
func isSymmetric2(root *TreeNode) bool {
	// 1. Both empty
	if root == nil {
		return true
	}
	// 2. Only one empty
	if root.Left == nil || root.Right == nil {
		return false
	}
	// 3. Both non-empty
	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		// 3.1. Pop 2 nodes from queue
		n1, n2 := queue[0], queue[1]
		queue = queue[2:]
		// 3.2. Check if they are the same
		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil || n1.Val != n2.Val {
			return false
		}
		// 3.3. Add left and right children to queue
		queue = append(queue, n1.Left, n2.Right)
		queue = append(queue, n1.Right, n2.Left)
	}
	return true
}
