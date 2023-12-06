// https://leetcode.com/problems/3sum/
package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := threeSum2(nums)
	for _, triplet := range result {
		fmt.Printf("%v\n", triplet)
	}
}

// solution 2: sort and two pointers
// O(n^2)
func threeSum2(num []int) [][]int {
	var result [][]int
	sort.Ints(num)
	for i := 0; i < len(num)-2; i++ {
		if i > 0 && num[i] == num[i-1] {
			continue
		}
		j := i + 1
		k := len(num) - 1
		for j < k {
			if num[i]+num[j]+num[k] == 0 {
				result = append(result, []int{num[i], num[j], num[k]})
				j++
				k--
				for j < k && num[j] == num[j-1] {
					j++
				}
				for j < k && num[k] == num[k+1] {
					k--
				}
			} else if num[i]+num[j]+num[k] < 0 {
				j++
			} else {
				k--
			}
		}
	}
	return result
}

// solution 3: hash table
// O(n^2)
func threeSum3(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums) // Sắp xếp mảng để dễ dàng kiểm tra trùng lặp

	for i := 0; i < len(nums)-2; i++ {
		// Kiểm tra trùng lặp cho i
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// Tạo map để lưu các giá trị đã xét
		m := make(map[int]int)
		for j := i + 1; j < len(nums); j++ {
			// Kiểm tra trùng lặp cho j
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			// Tìm giá trị cần thiết để tổng bằng 0
			c := 0 - nums[i] - nums[j]
			// Kiểm tra giá trị đó đã được xét chưa
			if _, ok := m[c]; ok {
				result = append(result, []int{nums[i], nums[j], c})
			} else {
				// Nếu chưa thì lưu lại giá trị j
				m[nums[j]] = j
			}
		}
	}
	return result
}

// solution 1: brute force
// O(n^3)
func threeSum(nums []int) [][]int {
	var result [][]int
	sort.Ints(nums) // Sắp xếp mảng để dễ dàng kiểm tra trùng lặp

	for i := 0; i < len(nums)-2; i++ {
		// Kiểm tra trùng lặp cho i
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		for j := i + 1; j < len(nums)-1; j++ {
			// Kiểm tra trùng lặp cho j
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < len(nums); k++ {
				// Kiểm tra trùng lặp cho k
				if k > j+1 && nums[k] == nums[k-1] {
					continue
				}
				if nums[i]+nums[j]+nums[k] == 0 {
					result = append(result, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}
	return result
}
