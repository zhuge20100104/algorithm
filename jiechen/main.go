package main

import (
	"fmt"
)

// Jiechen 求一个整数的阶乘的函数
func Jiechen(n int) int {
	if n == 1 {
		return 1
	}
	return n * Jiechen(n-1)
}

func main() {
	n := 3
	res := Jiechen(n)
	fmt.Printf("%d的阶乘是%d\n", n, res)
}
