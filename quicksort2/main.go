package main

import (
	"fmt"

	"github.com/zhuge20100104/algorithm/quicksort2/utils"
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
	QuickSort(arr, left, p[0]-1)
	// 递归排序右分区
	QuickSort(arr, p[1]+1, right)
}

// partion 三分法分区方法
// 总结:
// 小于主元，sp位置和eq位置交换，eq++；sp++
// 等于主元，sp++
// 大于主元，sp和r交换，r–；
// 最后的主元就是小于的最后一个位置，和大于的第一个位置，是一个数组[p1,p2]
func partition(arr []int, left, right int) []int {
	// 获取中轴值
	pivot := arr[left]

	// 等于分区指针
	eq := left

	for left <= right {
		if left <= right && arr[left] < pivot {
			arr[left], arr[eq] = arr[eq], arr[left]
			eq++
			left++
		} else if left <= right && arr[left] == pivot {
			left++
		} else if left <= right && arr[left] > pivot {
			arr[left], arr[right] = arr[right], arr[left]
			right--
		}
	}
	ret := []int{eq, right}
	return ret
}

// 三分法，有相同元素值的快速排序
func main() {
	arr := utils.U.GenerateRandomArr(10, 1, 100)
	fmt.Println("排序之前: ", arr)
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println("排序之后: ", arr)
}
