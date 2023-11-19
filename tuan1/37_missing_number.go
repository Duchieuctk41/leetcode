// https://leetcode.com/problems/missing-number/
package main

func main() {
	nums := []int{3, 0, 1, 4}
	println(missingNumber3(nums))
}

// solution 1: sort
func missingNumber1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 1. Sort the array
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[i] {
				nums[j], nums[i] = nums[i], nums[j]
			}
		}
	}
	// 2. Check if the first element is 0
	if nums[0] != 0 {
		return 0
	}
	// 3. Check if the last element is n
	if nums[len(nums)-1] != len(nums) {
		return len(nums)
	}
	// 4. Check if the difference between 2 consecutive elements is 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1]-nums[i] != 1 {
			return nums[i] + 1
		}
	}
	// 5. Return the missing number
	return 0
}

// solution 2: hash map
func missingNumber2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numMap := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		numMap[nums[i]] = true
	}
	for i := 0; i <= len(nums); i++ {
		if _, ok := numMap[i]; !ok {
			return i
		}
	}
	return 0
}

// solution 3: math
func missingNumber3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	n := len(nums)
	sum := n * (n + 1) / 2
	for i := 0; i < len(nums); i++ {
		sum -= nums[i]
	}
	return sum
}

// solution 4: bit manipulation
func missingNumber4(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	missing := len(nums)
	for i := 0; i < len(nums); i++ {
		missing ^= i ^ nums[i]
	}
	return missing
}

func missingNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	numMap := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if _, ok := numMap[nums[i]]; ok {
			delete(numMap, nums[i])
		} else {
			if nums[i] != 0 {
				numMap[nums[i]] = true
			}
		}
		if _, ok := numMap[i+1]; ok {
			delete(numMap, i+1)
		} else {
			numMap[i+1] = true
		}
	}
	for k, _ := range numMap {
		return k
	}
	return 0
}
