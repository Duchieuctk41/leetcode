// https://leetcode.com/problems/middle-of-the-linked-list/
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}

func middleNode(head *ListNode) *ListNode {
	fast := head
	slow := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	return slow
}
