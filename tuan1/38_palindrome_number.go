// https://leetcode.com/problems/palindrome-number/
package main

import "strconv"

func main() {
	println(isPalindrome2(121))
}

// solution 1: convert to string
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	// 1. Convert the number to a string
	str := strconv.Itoa(x)
	// 2. Check if the string is a palindrome
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}
	return true
}

// solution 2: math
func isPalindrome2(x int) bool {
	if x < 0 {
		return false
	}
	// 1. Reverse the number
	rev := 0
	tmp := x
	for tmp > 0 {
		rev = rev*10 + tmp%10
		tmp /= 10
	}
	// 2. Compare the number with the reversed number
	return x == rev
}
