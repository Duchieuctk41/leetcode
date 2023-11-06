// https://leetcode.com/problems/longest-palindrome/
package main

func longestPalindrome(s string) int {
	m := make(map[rune]int)
	for _, c := range s {
		m[c]++
	}
	res := 0
	odd := false
	for _, v := range m {
		if v%2 == 0 {
			res += v
		} else {
			res += v - 1
			odd = true
		}
	}
	if odd {
		res++
	}
	return res
}
