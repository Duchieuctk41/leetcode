// https://leetcode.com/problems/linked-list-cycle/
// use Floyd's Cycle Detection Algorithm
package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	// case 1: true (3 -> 2 -> 0 -> -4 -> 2)
	head := &ListNode{Val: 3,
		Next: &ListNode{Val: 2,
			Next: &ListNode{Val: 0,
				Next: &ListNode{Val: -4}}}}
	head.Next.Next.Next.Next = head.Next

	fmt.Print(hasCycle(head))
}

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}
