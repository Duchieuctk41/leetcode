// https://leetcode.com/problems/counting-bits/
package main

func main() {
	println(countBits(2))
	println(countBits(5))
}

// solution 1: brute force
func countBits1(num int) []int {
	result := make([]int, num+1)
	for i := 0; i <= num; i++ {
		result[i] = countOne(i)
	}
	return result
}

func countOne(num int) int {
	count := 0
	for num > 0 {
		num &= num - 1
		count++
	}
	return count
}

// solution 2: dynamic programming
// f[i] = f[i/2] + i%2
func countBits2(num int) []int {
	result := make([]int, num+1)
	for i := 1; i <= num; i++ {
		result[i] = result[i/2] + i%2
	}
	return result
}

// solution 3: dynamic programming
// f[i] = f[i>>1] + i&1
func countBits(num int) []int {
	result := make([]int, num+1)
	for i := 1; i <= num; i++ {
		result[i] = result[i>>1] + i&1
	}
	return result
}
