package main

import (
	"fmt"
	"time"

	"github.com/zhuge20100104/algorithm/poweran/utils"
)

const (
	// PowTimes 幂运算的次数
	PowTimes = 60
)

// Pow0 求a的n次幂的暴力解法
// 时间复杂度为 N
func Pow0(a int, n int) int {
	res := 1
	for i := 0; i < n; i++ {
		res *= a
	}
	return res
}

// Pow1 使用递归，拆一半，优化到lgN
func Pow1(a int, n int) int {
	// 递归退出标识
	if n == 0 {
		return 1
	}

	// 用于翻番的指数
	exp := 1
	res := a

	// 还可以接着翻番
	for (exp << 1) <= n {
		res *= res
		exp <<= 1
	}
	// 可能还有没算完的，递归进行计算
	return res * Pow1(a, n-exp)
}

func main() {
	// 打印传统算法耗时
	start := time.Now()
	fmt.Println(Pow0(2, PowTimes))
	utils.U.PrintTimeDiffMS(&start)

	// 打印二分算法耗时
	start = time.Now()
	fmt.Println(Pow1(2, PowTimes))
	utils.U.PrintTimeDiffMS(&start)
}
