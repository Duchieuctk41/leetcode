// https://leetcode.com/problems/palindrome-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// solution 1: use stack
func isPalindrome(head *ListNode) bool {
	stack := make([]int, 0)
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	for i := 0; i < len(stack)/2; i++ {
		if stack[i] != stack[len(stack)-1-i] {
			return false
		}
	}
	return true
}

// solution 2: use two pointers
// 1. find the middle of the linked list
// 2. reverse the second half of the linked list
// 3. compare the first half and the second half
func isPalindrome2(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	// find the middle of the linked list
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
	}
	// reverse the second half of the linked list
	var prev *ListNode
	for slow != nil {
		slow.Next, prev, slow = prev, slow, slow.Next
	}
	// compare the first half and the second half
	for prev != nil {
		if head.Val != prev.Val {
			return false
		}
		head, prev = head.Next, prev.Next
	}
	return true
}
