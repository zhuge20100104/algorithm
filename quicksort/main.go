package main

import (
	"fmt"

	"github.com/zhuge20100104/algorithm/quicksort/utils"
)

// QuickSort 快速排序算法
// arr 数组
// left 左指针
// right 右指针
func QuickSort(arr []int, left int, right int) {
	if left > right {
		return
	}
	p := partition(arr, left, right)
	QuickSort(arr, left, p-1)
	QuickSort(arr, p+1, right)
}

func partition(arr []int, left int, right int) int {
	// 获取中轴值
	pivot := arr[left]
	// 保存中轴值所在的位置
	l := left
	// 当左右节点不重合时，做左右的划分排序
	for left <= right {
		// 如果左边元素小于中轴值，左指针直接往右移
		if left <= right && arr[left] <= pivot {
			left++
			// 如果左边元素大于中轴值，说明它应该被放在右边，左右互换，右边元素--
		} else if left <= right && arr[left] > pivot {
			// 左右互换，此时右侧多了一个大于pivot的元素
			arr[left], arr[right] = arr[right], arr[left]
			right--
		}
	}

	arr[l], arr[right] = arr[right], arr[l]
	return right
}

// 单向扫描法的快速排序
func main() {
	// 使用工具类生成一个随机切片
	arr := utils.U.GenerateRandomArr(30, 1, 100)
	fmt.Println("排序之前:", arr)
	// 排序
	QuickSort(arr, 0, len(arr)-1)
	// 打印结果
	fmt.Println("排序之后:", arr)

}
