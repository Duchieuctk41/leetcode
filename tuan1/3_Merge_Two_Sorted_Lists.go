package main

import "fmt"

func main() {
	var list1 = []int{1, 2, 3}
	var list2 = []int{1, 2, 4}
	var result = mergeTwoLists(list1, list2)
	fmt.Println(result)
}

func mergeTwoLists(list1, list2 []int) []int {
	var result []int
	i, j := 0, 0

	for i < len(list1) && j < len(list2) {
		if list1[i] < list2[j] {
			result = append(result, list1[i])
			i++
		} else {
			result = append(result, list2[j])
			j++
		}
	}

	for i < len(list1) {
		result = append(result, list1[i])
		i++
	}

	for j < len(list2) {
		result = append(result, list2[j])
		j++
	}

	return result
}

// ----------------- 2 -----------------

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var result, tmp *ListNode
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val <= l2.Val {
		tmp = l1
		l1 = l1.Next
	} else {
		tmp = l2
		l2 = l2.Next
	}
	result = tmp
	for ; l1 != nil; l1 = l1.Next {
		for ; l2 != nil && l2.Val < l1.Val; l2 = l2.Next {
			tmp.Next = l2
			tmp = tmp.Next
		}
		tmp.Next = l1
		tmp = tmp.Next
	}
	if l2 != nil {
		tmp.Next = l2
	} else {
		tmp.Next = nil
	}
	return result
}
