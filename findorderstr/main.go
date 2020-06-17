package main

import (
	"fmt"
	"strings"
)

// FindOrderStr 在有空值的字符串数组中查找指定字符串
// arr 被查找的字符串数组
// key 要查找的字符串
func FindOrderStr(arr []string, key string) int {
	begin := 0
	end := len(arr) - 1

	for begin <= end {
		mid := begin + ((end - begin) >> 1)
		// 遇到了空串，往一个方向查找，直到不是空串
		for arr[mid] == "" {
			mid++
			// 已经越界，说明找不着， return -1
			if mid > end {
				return -1
			}
		}

		if strings.Compare(key, arr[mid]) < 0 {
			end = mid - 1
		} else if strings.Compare(key, arr[mid]) > 0 {
			begin = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

func main() {
	arr := []string{"a", "", "ab", "", "abc", "b", "", "ba", "bac"}
	index := FindOrderStr(arr, "dac")
	fmt.Printf("Index的值为: %v\n", index)
}
