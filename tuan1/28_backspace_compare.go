// https://leetcode.com/problems/backspace-string-compare/
package main

import "fmt"

func main() {
	s := "ab#c"
	t := "a#cd"
	fmt.Println(backspaceCompare(s, t))
}

// solution 1: two pointers
// Time complexity: O(n)
func backspaceCompare1(s string, t string) bool {
	i, j := len(s)-1, len(t)-1
	for {
		i = next(s, i)
		j = next(t, j)
		if i < 0 || j < 0 || s[i] != t[j] {
			break
		}
		i--
		j--
	}
	return i == -1 && j == -1
}

func next(s string, i int) int {
	back := 0
	for i >= 0 {
		if s[i] == '#' {
			back++
		} else if back > 0 {
			back--
		} else {
			break
		}
		i--
	}
	return i
}

// solution 2: build string
// Time complexity: O(n)
func backspaceCompare(s string, t string) bool {
	return build(s) == build(t)
}

func build(s string) string {
	var stack []rune
	for _, c := range s {
		if c != '#' {
			stack = append(stack, c)
		} else if len(stack) > 0 {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}
