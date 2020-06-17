package main

import (
	"fmt"

	"github.com/zhuge20100104/algorithm/quicksort1/utils"
)

// QuickSort 双向扫描的快速排序算法
// arr 待排序的数组
// left 左指针
// right 右指针
func QuickSort(arr []int, left, right int) {
	if left > right {
		return
	}
	// 分区
	p := partition(arr, left, right)
	// 递归排序左分区
	QuickSort(arr, left, p-1)
	// 递归排序右分区
	QuickSort(arr, p+1, right)
}

func partition(arr []int, left, right int) int {
	pivot := arr[left]

	// 记录中轴值所在位置
	l := left

	for left <= right {
		for left <= right && arr[left] <= pivot {
			left++
		}

		for left <= right && arr[right] > pivot {
			right--
		}

		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
		}
	}
	// 中轴值和right位置互换，把中轴值放到合理的位置
	arr[l], arr[right] = arr[right], arr[l]
	return right
}

// 基于 双向扫描法 的快速排序
func main() {
	arr := utils.U.GenerateRandomArr(10, 1, 100)
	fmt.Println("排序之前: ", arr)
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println("排序之后: ", arr)
}
