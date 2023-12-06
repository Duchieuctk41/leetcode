// https://leetcode.com/problems/implement-queue-using-stacks/
package main

import "fmt"

type MyQueue struct {
	data []int
}

func Constructor() MyQueue {
	return MyQueue{
		data: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.data = append(this.data, x)
}

func (this *MyQueue) Pop() int {
	if len(this.data) == 0 {
		return -1
	}
	res := this.data[0]
	this.data = this.data[1:]
	return res
}

func (this *MyQueue) Peek() int {
	if len(this.data) == 0 {
		return -1
	}
	return this.data[0]
}

func (this *MyQueue) Empty() bool {
	return len(this.data) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

//==============================================================================

// Queue : data structure
type Queue struct {
	data []int
}

// New : returns a new instance of a Queue
func New() *Queue {
	return &Queue{
		data: []int{},
	}
}

// IsEmpty : checks whether the queue is empty
func (q *Queue) IsEmpty() bool {
	return len(q.data) == 0
}

// Peek : return the next element in the queue
func (q *Queue) Peek() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queue is empty")
	}
	return q.data[0], nil
}

// Enqueue : adds an element onto the queue
func (q *Queue) Enqueue(n int) *Queue {
	q.data = append(q.data, n)
	return q
}

// Dequeue : removes the next element from the queue
// and returns its value
func (q *Queue) Dequeue() (int, error) {
	if len(q.data) == 0 {
		return 0, fmt.Errorf("Queue is empty")
	}
	element := q.data[0]
	q.data = q.data[1:]
	return element, nil
}