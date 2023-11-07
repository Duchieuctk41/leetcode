// https://leetcode.com/problems/roman-to-integer/
package main

func main() {
	romanToInt("MCMXCIV")
}

func romanToInt(s string) int {
	romanMap := map[string]int{
		"I": 1, "V": 5, "X": 10, "L": 50,
		"C": 100, "D": 500, "M": 1000,
	}
	result := 0
	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && romanMap[string(s[i])] < romanMap[string(s[i+1])] {
			result -= romanMap[string(s[i])]
		} else {
			result += romanMap[string(s[i])]
		}
	}
	return result
}
