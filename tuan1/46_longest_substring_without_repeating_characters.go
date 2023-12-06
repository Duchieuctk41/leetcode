// https://leetcode.com/problems/longest-substring-without-repeating-characters/
package main

// solution 1: brute force
func lengthOfLongestSubstring(s string) int {
	var max int
	for i := 0; i < len(s); i++ {
		var m = make(map[byte]bool)
		for j := i; j < len(s); j++ {
			if _, ok := m[s[j]]; ok {
				break
			}
			m[s[j]] = true
		}
		if len(m) > max {
			max = len(m)
		}
	}
	return max
}

// solution 2: sliding window
func lengthOfLongestSubstring2(s string) int {
	var max int
	var m = make(map[byte]int)
	for i, j := 0, 0; j < len(s); j++ {
		if _, ok := m[s[j]]; ok {
			i = maxInt(m[s[j]], i)
		}
		max = maxInt(max, j-i+1)
		m[s[j]] = j + 1
	}
	return max
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// solution 3: sliding window optimized
func lengthOfLongestSubstring3(s string) int {
	var max int
	var m = make(map[byte]int)
	for i, j := 0, 0; j < len(s); j++ {
		if _, ok := m[s[j]]; ok {
			i = maxInt(m[s[j]], i)
		}
		max = maxInt(max, j-i+1)
		m[s[j]] = j + 1
	}
	return max
}

// solution 4: sliding window optimized with hash map

func lengthOfLongestSubstring4(s string) int {
	charIndexMap := make(map[rune]int)
	maxLength := 0
	start := 0

	for i, char := range s {
		if lastIndex, found := charIndexMap[char]; found && lastIndex >= start {
			start = lastIndex + 1
		}
		charIndexMap[char] = i
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
	}

	return maxLength
}
