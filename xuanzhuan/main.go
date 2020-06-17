package main

import (
	"fmt"
)

// Min 使用二分法找出旋转后数组的最小值
func Min(arr []int) int {
	left := 0
	right := len(arr) - 1
	// 数组本身是有序的
	if arr[left] < arr[right] {
		return arr[left]
	}

	mid := left
	for left+1 < right {
		mid = left + ((right - left) >> 1)
		if arr[mid] < arr[left] {
			right = mid
		} else {
			left = mid
		}
	}
	return arr[right]
}

func main() {
	arr := []int{1, 2, 3, 4, 5}
	res := Min(arr)
	fmt.Printf("最小值为: %v\n", res)
}
