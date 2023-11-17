package main

func main() {
	println(singleNumber([]int{4, 1, 2, 1, 2}))
}

// solution 1: use map
// 1. create a map to store the number of occurrences of each number
// 2. if the number of occurrences of a number is 2, delete it from the map
// 3. return the only number in the map
func singleNumber(nums []int) int {
	numsMap := make(map[int]bool)
	for _, v := range nums {
		if _, ok := numsMap[v]; ok {
			delete(numsMap, v)
		} else {
			numsMap[v] = true
		}
	}
	for k := range numsMap {
		return k
	}
	return 0
}

// solution 2: use bit operation
// 1. XOR all numbers
// 2. return the result
func singleNumber2(nums []int) int {
	result := 0
	for _, v := range nums {
		result ^= v
	}
	return result
}
