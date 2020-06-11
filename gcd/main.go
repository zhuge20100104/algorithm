package main

import (
	"fmt"
)

// Gcd 求两数的最大公约数
func Gcd(m, n int) int {
	if m%n == 0 {
		return n
	}
	return Gcd(n, m%n)
}

func main() {
	m, n := 36, 8
	res := Gcd(m, n)
	fmt.Printf("%v和%v 的最大公约数是 %v\n", m, n, res)
}
