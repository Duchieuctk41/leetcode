// https://leetcode.com/problems/maximum-subarray/
package main

import "fmt"

func main() {
	nums := []int{-2, -1, -3, -4, -1, -2, -1, -5, -4}
	fmt.Println(maxSubArray(nums))

}

func maxSubArray(nums []int) int {
	max := nums[0]
	sum := 0
	for _, v := range nums {
		if sum < 0 {
			sum = 0
		}
		sum += v
		if sum > max {
			max = sum
		}
	}
	return max
}
