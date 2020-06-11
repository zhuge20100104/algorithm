package main

import (
	"fmt"
)

// ReverseStr 反转字符串的函数
func ReverseStr(str string, k int) string {
	if k == 0 {
		return str[0:1]
	}
	return str[k:k+1] + ReverseStr(str, k-1)
}

func main() {
	str := "abcd"
	res := ReverseStr(str, len(str)-1)
	fmt.Printf("%v反转的结果为: %v\n", str, res)
}
