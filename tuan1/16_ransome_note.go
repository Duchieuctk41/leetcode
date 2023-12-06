// https://leetcode.com/problems/ransom-note/
package main

// solution: hash table
func canConstruct(ransomNote string, magazine string) bool {
	m := make(map[rune]int)
	for _, v := range magazine {
		m[v]++
	}
	for _, v := range ransomNote {
		if m[v] == 0 {
			return false
		}
		m[v]--
	}
	return true
}


func canConstruct(ransomNote string, magazine string) bool {
	byteIndex := make([]int, 26)
	for _, letter := range magazine {
		byteIndex[int(letter - 'a')]++
	}
	for _, ch := range ransomNote {
		byteIndex[int(ch - 'a')]--
	}

	for i := range byteIndex {
		if byteIndex[i] < 0 {
			return false
		}
	}
	return true
}