package main

import (
	"fmt"

	"github.com/zhuge20100104/algorithm/quicksort3/utils"
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

// getMidValueIdx 获取中值索引的函数
func getMidValueIdx(arr []int, left int, right int) int {
	// 数组中间元素的索引
	midIndex := left + ((right - left) >> 1)
	// 居中值所在的实际索引
	midIndexVal := -1
	if (arr[left] >= arr[right] && arr[left] <= arr[midIndex]) ||
		(arr[left] >= arr[midIndex] && arr[left] <= arr[right]) {
		midIndexVal = left
	} else if (arr[right] >= arr[left] && arr[right] <= arr[midIndex]) ||
		(arr[right] >= arr[midIndex] && arr[right] <= arr[left]) {
		midIndexVal = right
	} else {
		midIndexVal = midIndex
	}
	return midIndexVal
}

func partition(arr []int, left int, right int) int {

	// 获取中间值所在的索引
	midIndexVal := getMidValueIdx(arr, left, right)
	// 最左边元素和中间值元素互换
	arr[left], arr[midIndexVal] = arr[midIndexVal], arr[left]

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

func main() {
	arr := utils.U.GenerateRandomArr(10, 1, 100)
	fmt.Println("排序前,", arr)
	QuickSort(arr, 0, len(arr)-1)
	fmt.Println("排序后,", arr)
}
