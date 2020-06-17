package main

import (
	"fmt"
	"time"

	"github.com/zhuge20100104/algorithm/compareper/utils"
)

const (
	// ArrayLen 大数组长度
	ArrayLen = 10000 * 10000
)

// SeqSearch 顺序查找算法
// 顺序查找算法时间复杂度 N
func SeqSearch(arr []int, ele int) int {
	for i := 0; i < len(arr); i++ {
		if ele == arr[i] {
			return i
		}
	}
	return -1
}

// BinSearch 二分查找算法
// 二分查找时间复杂度 N*lgN
func BinSearch(arr []int, low, high, key int) int {
	for low <= high {
		mid := low + ((high - low) >> 1)
		midVal := arr[mid]
		if key > midVal {
			low = mid + 1
		} else if key < midVal {
			high = mid - 1
		} else {
			return mid
		}
	}
	return -1
}

// 顺序查找与二分查找的性能比较
func main() {
	arr := make([]int, 0, ArrayLen)
	for i := 0; i < ArrayLen; i++ {
		arr = append(arr, i+1)
	}

	// 查找最后一个元素
	target := ArrayLen

	// 二分查找性能
	start := time.Now()
	index := BinSearch(arr, 0, len(arr)-1, target)
	utils.U.PrintTimeDiffMS(&start)
	fmt.Println("元素索引为", index)

	// 顺序查找性能
	start = time.Now()
	index = SeqSearch(arr, target)
	utils.U.PrintTimeDiffMS(&start)
	fmt.Println("元素索引为", index)
}
