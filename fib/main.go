package main

import (
	"fmt"
)

// Fib 计算斐波那契数列的函数
func Fib(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

func main() {
	n := 5
	res := Fib(n)
	fmt.Printf("第 %v 个斐波那契数列是 %v\n", n, res)
}
