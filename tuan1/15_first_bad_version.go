// https://leetcode.com/problems/first-bad-version/

package main

import "fmt"

func main() {
	fmt.Println(firstBadVersion(5))
}

// solution: binary search
func firstBadVersion(n int) int {
	l, r := 1, n
	for l < r {
		mid := (l + r) / 2
		if isBadVersion(mid) {
			r = mid
		} else {
			l = mid + 1
		}
	}

	return l
}

// Mock function for testing
func isBadVersion(version int) bool {
	if version >= 4 {
		return true
	}
	return false
}
