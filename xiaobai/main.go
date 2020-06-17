package main

import (
	"fmt"
)

// UpLift 计算小白上楼梯有几种上法
// 时间复杂度为 n 的3次方
func UpLift(n int) int {
	if n == 0 {
		return 1
	}

	if n == 1 {
		return 1
	}

	if n == 2 {
		return 2
	}
	return UpLift(n-1) + UpLift(n-2) + UpLift(n-3)
}

func main() {
	for {
		fmt.Printf("请输入需要上的楼梯数:")
		var steps int
		fmt.Scanln(&steps)
		ways := UpLift(steps)
		fmt.Printf("总共有 %v 种上楼梯的方法\n", ways)
	}
}
