package main

import "fmt"

// LongestSub 求最长递增子序列
func LongestSub(arr []int) int {
	begin := 0
	// end 要把begin算进去，所以要加1
	end := 1
	maxLen := 0
	for i := 0; i < len(arr); i++ {
		if i < len(arr)-1 && arr[i] <= arr[i+1] {
			end++
		} else {
			if end-begin > maxLen {
				maxLen = end - begin
			}

			begin = i
			// end 要把begin算进去，所以要加1
			end = begin + 1
		}
	}
	return maxLen
}

func main() {
	arr := []int{1, 9, 2, 5, 7, 9, 10, 3, 4, 6, 8, 0}
	res := LongestSub(arr)
	fmt.Printf("最长子串长度为: %v\n", res)
}
