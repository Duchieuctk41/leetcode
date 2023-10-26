// https://leetcode.com/problems/valid-palindrome/
package main

func main() {

}

func isAlphaNumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') ||
		(c >= '0' && c <= '9') ||
		(c >= 'A' && c <= 'Z')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + 32
	}

	return c
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}

	i, j := 0, len(s)-1 // i: left, j: right
	for i < j {
		if !isAlphaNumeric(s[i]) {
			i++
			continue
		}

		if !isAlphaNumeric(s[j]) {
			j--
			continue
		}

		if toLower(s[i]) != toLower(s[j]) {
			return false
		}

		i++
		j--
	}

	return true
}
