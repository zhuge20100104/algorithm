package main

import (
	"fmt"
)

// Hannota 汉诺塔打印步骤的函数
func Hannota(N int, from, to, help string) {
	if N == 1 {
		fmt.Printf("将第 %v 个盘从 %v 移动到 %v\n", N, from, to)
		return
	}

	Hannota(N-1, from, help, to)
	fmt.Printf("将第 %v 个盘从 %v 移动到 %v\n", N, from, to)
	Hannota(N-1, help, to, from)
}

func main() {
	Hannota(3, "A", "B", "C")
}
