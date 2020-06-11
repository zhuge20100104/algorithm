package main

import (
	"fmt"
)

// BinSearch 二分查找方法
// 返回查找的 key 在数组中的位置
func BinSearch(arr []int, left, right, key int) int {
	// 说明没找着, return -1
	if left > right {
		return -1
	}

	mid := (left + right) / 2
	// 说明元素在 mid 右边，在mid右边继续查找
	if key > arr[mid] {
		return BinSearch(arr, mid+1, right, key)
		// 说明在 mid 左边,在mid左边继续查找
	} else if key < arr[mid] {
		return BinSearch(arr, left, mid-1, key)
		//找着了，中码
	} else {
		return mid
	}
}

func main() {
	arr := []int{1, 3, 6, 8, 10, 18, 21, 25, 30, 33, 37}
	key := 32
	res := BinSearch(arr, 0, len(arr)-1, key)
	fmt.Printf("%v在arr中的位置是 :%v\n", key, res)
}
