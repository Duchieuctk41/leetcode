// https://leetcode.com/problems/two-sum/

package main


// solution 1: brute force (O(n^2)
func twoSum(nums []int, target int) []int {
	for i,v := range nums{
			for j:=i+1;j<len(nums);j++{
					if nums[j]==target-v{
					return []int{i,j}
					}
			}
	}
	panic("should never happen")
}

// solution 2: hash table (O(n))
func twoSum(nums []int, target int) []int {
	indexMap := make(map[int]int)
	for currIndex, currentValue := range nums {
		if requiredIdx, ok := indexMap[target-currentValue]; ok {
			return []int{requiredIdx, currIndex}
		}
		indexMap[currentValue] = currIndex
	}
	return []int{}
}