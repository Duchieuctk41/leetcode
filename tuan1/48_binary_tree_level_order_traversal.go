// https://leetcode.com/problems/binary-tree-level-order-traversal/
package main

// solution 1: BFS
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var level []int
		var nextQueue []*TreeNode
		for _, node := range queue {
			level = append(level, node.Val)
			if node.Left != nil {
				nextQueue = append(nextQueue, node.Left)
			}
			if node.Right != nil {
				nextQueue = append(nextQueue, node.Right)
			}
		}
		res = append(res, level)
		queue = nextQueue
	}
	return res
}

// solution 2: DFS
func levelOrder2(root *TreeNode) [][]int {
	var res [][]int
	dfs(root, 0, &res)
	return res
}

func dfs(root *TreeNode, level int, res *[][]int) {
	if root == nil {
		return
	}
	if len(*res) == level {
		*res = append(*res, []int{})
	}
	(*res)[level] = append((*res)[level], root.Val)
	dfs(root.Left, level+1, res)
	dfs(root.Right, level+1, res)
}
