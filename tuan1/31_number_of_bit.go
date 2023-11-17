// https://leetcode.com/problems/number-of-1-bits/
package main

func main() {
	println(hammingWeight(00000000000000000000000000001011))
	println(hammingWeight(00000000000000000000000010000000))
}

// solution 1: use bit operation
// 1. check if the last bit is 1
// 2. shift right 1 bit
// 3. repeat 1, 2 until num == 0
func hammingWeight(num uint32) int {
	count := 0
	for num > 0 {
		if num&1 == 1 {
			count++
		}
		num >>= 1
	}
	return count
}
