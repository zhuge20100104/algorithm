package main

import (
	"fmt"
)

// SherSort 希尔排序算法实现
// 分组插入排序，主键缩小分组范围
func SherSort(arr []int) {
	for interval := len(arr) / 2; interval > 0; interval /= 2 {
		for i := interval; i < len(arr); i++ {
			val := arr[i]
			index := i - interval
			for index > -1 && val < arr[index] {
				arr[index+interval] = arr[index]
				index -= interval
			}
			arr[index+interval] = val
		}
	}
}

func main() {
	arr := []int{9, 8, 7, 6, 5, 6, 3, 2, 1}
	SherSort(arr)
	fmt.Println(arr)
}
