// https://leetcode.com/problems/add-binary/

package main

func main() {
	println(addBinary("11", "1"))
	println(addBinary("1010", "1011"))
}

func addBinary(a string, b string) string {
	var res string
	var carry int
	var i, j = len(a) - 1, len(b) - 1

	for i >= 0 || j >= 0 || carry > 0 {
		if i >= 0 {
			carry += int(a[i] - '0')
			i--
		}

		if j >= 0 {
			carry += int(b[j] - '0')
			j--
		}

		res = string(carry%2+'0') + res
		carry /= 2
	}

	return res
}
