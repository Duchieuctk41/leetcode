// https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	var (
		pVal = p.Val
		qVal = q.Val
	)
	for root != nil {
		if pVal > root.Val && qVal > root.Val {
			root = root.Right
		} else if pVal < root.Val && qVal < root.Val {
			root = root.Left
		} else {
			return root
		}
	}
	return root
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root.Val > p.Val && root.Val > q.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if root.Val < p.Val && root.Val < q.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	// root.Val == p.Val || root.Val == q.Val || (root.Val > p.Val && root.Val < q.Val) || (root.Val < p.Val && root.Val > q.Val)
	return root
}
