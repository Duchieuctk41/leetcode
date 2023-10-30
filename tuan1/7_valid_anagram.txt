// https://leetcode.com/problems/valid-anagram/
package main

func main() {
	s := "anaggam"
	t := "nagaram"
	println(isAnagram(s, t))
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// count char in s
	count := make(map[rune]int)
	for _, c := range s {
		count[c]++
	}

	// count char in t
	for _, c := range t {
		count[c]--
	}

	// check count
	for _, v := range count {
		if v != 0 {
			return false
		}
	}

	return true
}
