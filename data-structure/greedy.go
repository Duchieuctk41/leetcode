package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

func findMinCoins(coins []int, amount int) []int {
	// Sắp xếp mảng coins giảm dần
	sort.Sort(sort.Reverse(sort.IntSlice(coins)))

	result := make([]int, 0)
	for _, coin := range coins {
		// Trong khi số tiền còn lại lớn hơn giá trị của coin hiện tại
		for amount >= coin {
			// Trừ coin hiện tại khỏi amount
			amount -= coin
			// Thêm coin vào kết quả
			result = append(result, coin)
		}
		// Nếu đã trả đủ tiền, thoát khỏi vòng lặp
		if amount == 0 {
			break
		}
	}

	// Trả về kết quả nếu số tiền đã được trả đủ, ngược lại trả về nil
	if amount == 0 {
		return result
	} else {
		return nil
	}
}

func main() {
	tst()
}

func tst() {
	// Tạo chuỗi có độ dài 100 ký tự, chứa ký tự "a" ở vị trí cuối cùng
	longString := strings.Repeat("w_acc_ass,", 4)
	longString += "w_acc_men"

	// Đo thời gian tìm kiếm
	startTime := time.Now()
	result := strings.Contains(longString, "w_acc_men")
	elapsedTime := time.Since(startTime)

	// In kết quả và thời gian tìm kiếm
	fmt.Printf("Chuỗi có chứa ký tự 'w_acc_men': %t\n", result)
	fmt.Printf("Thời gian tìm kiếm: %s\n", elapsedTime)
}
