// https://leetcode.com/problems/move-zeroes/
package main

// solution 1: no swap
func moveZeroes(nums []int) {
	// 1. Count number of zeros
	// 2. Move non-zero elements to the left
	// 3. Fill the remaining array with zeros
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[count] = nums[i]
			count++
		}
	}
	for i := count; i < len(nums); i++ {
		nums[i] = 0
	}
}

// solution 2: swap
func moveZeroes2(nums []int) {
	// 1. Count number of zeros
	// 2. Move non-zero elements to the left
	// 3. Fill the remaining array with zeros
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[count], nums[i] = nums[i], nums[count]
			count++
		}
	}
}
