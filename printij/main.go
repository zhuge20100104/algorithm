package main

import (
	"fmt"
)

// PrintIj 打印i到j的函数
func PrintIj(i, j int) {
	// 如果已经打印到j了，则退出
	if i-1 == j {
		return
	}

	fmt.Println(i)
	PrintIj(i+1, j)
}

func main() {
	PrintIj(1, 10)
}
