// Xuất ra dãy số fibonacci có trong đoạn từ 0 - n

package main

import "fmt"

// fibonaci viết theo cách thông thường
func fibonacci_normal(n int) int {
	first, second := 0, 1
	for i := 0; i < n; i++ {
		ret := first
		first, second = second, first+second
		fmt.Println(ret)
	}
	return first
}

// fibonacci viết theo cách closure
func fibonacci_closure() func() int {
	first, second := 0, 1
	return func() int {
		ret := first
		first, second = second, first+second
		return ret
	}
}

func main() {
	fibonacci_normal(10)
	fib_closure := fibonacci_closure()
	for i := 0; i < 10; i++ {
		fmt.Println(fib_closure())
	}
}
