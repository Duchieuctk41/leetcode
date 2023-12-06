// https://leetcode.com/problems/squares-of-a-sorted-array/
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	fmt.Println(sortedSquares5(nums))
}

// solution 1: brute force
func sortedSquares(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] *= nums[i]
	}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

// solution 2: sort
func sortedSquares2(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		nums[i] *= nums[i]
	}
	sort.Ints(nums)
	return nums
}

// solution 3: two pointer
func sortedSquares3(nums []int) []int {
	res := make([]int, len(nums))
	l, r := 0, len(nums)-1
	for i := len(nums) - 1; i >= 0; i-- {
		if abs(nums[l]) > abs(nums[r]) {
			res[i] = nums[l] * nums[l]
			l++
		} else {
			res[i] = nums[r] * nums[r]
			r--
		}
	}
	return res
}

// lấy giá trị tuyệt đối
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

// solution 4: binary search
func sortedSquares4(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = nums[i] * nums[i]
	}
	sort.Ints(res)
	return res
}

// solution 5: best
func sortedSquares5(nums []int) []int {
	res := make([]int, len(nums))
	l, r := 0, len(nums)-1
	for i := len(nums) - 1; i >= 0; i-- {
		if abs(nums[l]) > abs(nums[r]) {
			res[i] = nums[l] * nums[l]
			l++
		} else {
			res[i] = nums[r] * nums[r]
			r--
		}
	}
	return res
}
