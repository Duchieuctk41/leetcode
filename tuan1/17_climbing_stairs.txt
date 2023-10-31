// https://leetcode.com/problems/climbing-stairs/
package main

// solution 1: recursion
func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	return climbStairs(n-1) + climbStairs(n-2)
}

func climbStairs(n int) int {
	if n == 1 {
		return 1
	}
	f1, f2 := 1, 2
	for i := 3; i <= n; i++ {
		f1, f2 = f2, f1+f2
	}
	return f2
}
