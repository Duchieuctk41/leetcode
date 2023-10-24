// https://leetcode.com/problems/valid-parentheses/
package main

import "fmt"

func main() {
	s := "{()[]}"
	fmt.Println(isValid(s))
}
func isValid(s string) bool {
	if len(s) == 0 || len(s)%2 == 1 {
		return false
	}

	mapS := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}

	stack := []rune{}

	for _, v := range s {
		if _, ok := mapS[v]; ok {
			stack = append(stack, v)
		} else {
			a := stack[len(stack)-1]
			if mapS[a] != v {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
