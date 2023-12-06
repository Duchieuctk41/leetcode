// https://leetcode.com/problems/insert-interval/
package main

import (
	"fmt"
	"sort"
)

func main() {
	intervals := [][]int{{1, 2}, {3, 4}, {6, 9}}
	newInterval := []int{2, 5}
	fmt.Println(insert2(intervals, newInterval))
}

// solution 1: brute force
func insert(intervals [][]int, newInterval []int) [][]int {
	var res [][]int
	i := 0
	for i < len(intervals) && intervals[i][1] < newInterval[0] {
		res = append(res, intervals[i])
		i++
	}
	for i < len(intervals) && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}
	res = append(res, newInterval)
	for i < len(intervals) {
		res = append(res, intervals[i])
		i++
	}
	return res
}

// solution 2: binary search
func insert2(intervals [][]int, newInterval []int) [][]int {
	n := len(intervals)
	i := sort.Search(n, func(i int) bool { return intervals[i][0] > newInterval[0] })
	j := sort.Search(n, func(j int) bool { return intervals[j][1] > newInterval[1] })
	if i >= 1 && newInterval[0] <= intervals[i-1][1] {
		newInterval[0] = intervals[i-1][0]
		i--
	}
	if j < n && newInterval[1] >= intervals[j][0] {
		newInterval[1] = intervals[j][1]
		j++
	}
	return append(intervals[:i], append([][]int{newInterval}, intervals[j:]...)...)
}
