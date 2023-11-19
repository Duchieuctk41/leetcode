// https://leetcode.com/problems/convert-sorted-array-to-binary-search-tree/
package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// solution 1: recursion
func sortedArrayToBST(nums []int) *TreeNode {

	if len(nums) == 0 {
		return nil
	}

	mid := len(nums) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = sortedArrayToBST(nums[:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])

	return root
}

// solution 2: recursion
func sortedArrayToBST2(nums []int) *TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	mid := left + (right-left)/2
	root := &TreeNode{Val: nums[mid]}
	root.Left = helper(nums, left, mid-1)
	root.Right = helper(nums, mid+1, right)

	return root
}

// solution 3: iteration
func sortedArrayToBST3(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

	root := &TreeNode{}
	stack := []*TreeNode{root}
	left := []int{0}
	right := []int{len(nums) - 1}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		l := left[len(left)-1]
		left = left[:len(left)-1]
		r := right[len(right)-1]
		right = right[:len(right)-1]

		mid := l + (r-l)/2
		node.Val = nums[mid]

		if l <= mid-1 {
			node.Left = &TreeNode{}
			stack = append(stack, node.Left)
			left = append(left, l)
			right = append(right, mid-1)
		}

		if mid+1 <= r {
			node.Right = &TreeNode{}
			stack = append(stack, node.Right)
			left = append(left, mid+1)
			right = append(right, r)
		}
	}

	return root
}
