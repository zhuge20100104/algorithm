package main

import (
	"fmt"
	"sort"
	"time"
)

const (
	// ArrayLen 数组长度对象
	ArrayLen = 10000
)

// BubbleSort 冒泡排序
// 时间复杂度为n方
func BubbleSort(arr []int) {
	len := len(arr)
	for i := 0; i < len; i++ {
		for j := 0; j < len-1-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

// MakeArray 创建一个新的Array切片并返回的函数
func MakeArray() []int {
	// 构造一个长度为ArrayLen的完全逆序的切片
	arr := make([]int, 0, ArrayLen)
	for i := 0; i < ArrayLen; i++ {
		arr = append(arr, ArrayLen-i)
	}
	return arr
}

func main() {

	arr := MakeArray()
	start := time.Now()
	BubbleSort(arr)
	duration := time.Since(start).Nanoseconds()
	fmt.Printf("冒泡排序耗时: %v 纳秒\n", duration)

	arr1 := MakeArray()
	start = time.Now()
	sort.Ints(arr1)
	duration = time.Since(start).Nanoseconds()
	fmt.Printf("快速排序耗时: %v 纳秒\n", duration)

}
