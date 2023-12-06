// https://leetcode.com/problems/k-closest-points-to-origin/
package main

import "sort"

// solution 1: sort
func kClosest(points [][]int, k int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		x2 := points[i][0]*points[i][0] + points[i][1]*points[i][1]
		y2 := points[j][0]*points[j][0] + points[j][1]*points[j][1]
		return x2 < y2
	})
	return points[:k]
}

// solution 2: quick select
func kClosest2(points [][]int, k int) [][]int {
	quickSelect(points, 0, len(points)-1, k)
	return points[:k]
}

func quickSelect(points [][]int, l, r, k int) {
	if l >= r {
		return
	}
	pivot := partition(points, l, r)
	if pivot == k {
		return
	} else if pivot < k {
		quickSelect(points, pivot+1, r, k)
	} else {
		quickSelect(points, l, pivot-1, k)
	}
}

func partition(points [][]int, l, r int) int {
	pivot := points[r]
	i := l
	for j := l; j < r; j++ {
		if compare(points[j], pivot) < 0 {
			points[i], points[j] = points[j], points[i]
			i++
		}
	}
	points[i], points[r] = points[r], points[i]
	return i
}

func compare(p1, p2 []int) int {
	return p1[0]*p1[0] + p1[1]*p1[1] - p2[0]*p2[0] - p2[1]*p2[1]
}

// solution 3: max heap
func kClosest3(points [][]int, k int) [][]int {
	h := &maxHeap{}
	for _, p := range points {
		h.Push(p)
		if h.Len() > k {
			h.Pop()
		}
	}
	res := make([][]int, 0, k)
	for h.Len() > 0 {
		res = append(res, h.Pop().([]int))
	}
	return res
}

type maxHeap [][]int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i][0]*h[i][0]+h[i][1]*h[i][1] > h[j][0]*h[j][0]+h[j][1]*h[j][1]
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.([]int))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
