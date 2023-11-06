// https://leetcode.com/problems/majority-element/
package main

// mang chi co 2 so khac nhau, tim so xuat hien nhieu nhat
func main() {
	nums := []int{2, 2, 1, 1, 1, 2}
	println(majorityElement(nums))
}

func majorityElement(nums []int) int {
	count := 0
	var candidate int
	for _, num := range nums {
		if count == 0 {
			candidate = num
		}
		if candidate == num {
			count++
		} else {
			count--
		}
	}
	return candidate
}
