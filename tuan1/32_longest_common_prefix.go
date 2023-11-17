// https://leetcode.com/problems/longest-common-prefix/
package main

func main() {
	println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
}

// solution 1: use isPrefix
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0] // prefix is the first string
	for i := 1; i < len(strs); i++ {
		// if prefix is not a prefix of strs[i]
		// then remove the last character of prefix
		for !isPrefix(prefix, strs[i]) {
			prefix = prefix[:len(prefix)-1]
		}
	}
	return prefix
}

// check if prefix is a prefix of str
func isPrefix(prefix, str string) bool {
	if len(prefix) > len(str) {
		return false
	}
	for i := 0; i < len(prefix); i++ {
		if prefix[i] != str[i] {
			return false
		}
	}
	return true
}
