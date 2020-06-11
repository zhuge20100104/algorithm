package main

import "fmt"

// InsertSort 插入排序的递归算法
func InsertSort(arr []int, k int) {
	if k == 0 {
		return
	}
	InsertSort(arr, k-1)
	val := arr[k]
	index := k - 1

	for index >= 0 && val < arr[index] {
		arr[index+1] = arr[index]
		index--
	}

	arr[index+1] = val
}

func main() {
	arr := []int{9, 7, 8, 6, 5, 3, 4, 2, 1, 6, 10, 60}
	InsertSort(arr, len(arr)-1)
	fmt.Println(arr)
}
