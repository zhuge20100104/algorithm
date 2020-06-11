package main

import "fmt"

// Sum 求数组的和
// arr 需要求和的数组
// n 递归时的索引，求第几个元素之后的和
func Sum(arr []int, k int) int {
	if k+1 == len(arr) {
		return arr[k]
	}
	return arr[k] + Sum(arr, k+1)
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	res := Sum(arr, 0)
	fmt.Printf("arr之和为%v\n", res)
}
