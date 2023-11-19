// https://leetcode.com/problems/reverse-bits/
package main

// solution 1: use bit manipulation
func reverseBits(num uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		result = result << 1
		result = result | (num & 1)
		num = num >> 1
	}
	return result
}

// solution 2: use divide and conquer
func reverseBits2(num uint32) uint32 {
	num = (num >> 16) | (num << 16)
	num = ((num & 0xff00ff00) >> 8) | ((num & 0x00ff00ff) << 8)
	num = ((num & 0xf0f0f0f0) >> 4) | ((num & 0x0f0f0f0f) << 4)
	num = ((num & 0xcccccccc) >> 2) | ((num & 0x33333333) << 2)
	num = ((num & 0xaaaaaaaa) >> 1) | ((num & 0x55555555) << 1)
	return num
}

// solution 3: map
func reverseBits3(num uint32) uint32 {
	var result uint32
	m := map[int]int{
		0:  31,
		1:  30,
		2:  29,
		3:  28,
		4:  27,
		5:  26,
		6:  25,
		7:  24,
		8:  23,
		9:  22,
		10: 21,
		11: 20,
		12: 19,
		13: 18,
		14: 17,
		15: 16,
		16: 15,
		17: 14,
		18: 13,
		19: 12,
		20: 11,
		21: 10,
		22: 9,
		23: 8,
		24: 7,
		25: 6,
		26: 5,
		27: 4,
		28: 3,
		29: 2,
		30: 1,
		31: 0,
	}
	for i := 0; i < 32; i++ {
		result = result | ((num & 1) << m[i])
		num = num >> 1
	}
	return result
}

// solution 4: math
func reverseBits4(num uint32) uint32 {
	var result uint32
	for i := 0; i < 32; i++ {
		result = result*2 + num%2
		num = num / 2
	}
	return result
}
