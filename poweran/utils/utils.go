package utils

import (
	"fmt"
	"time"
)

const (
	// NSToMSUnits 纳秒转到毫秒应该除以的常量
	NSToMSUnits = float64(1000 * 1000)
)

// Utils 暴露给外部使用的工具函数
type Utils struct{}

var (
	// U 暴露给外界使用的全局的Utils对象
	U *Utils
)

func init() {
	// 初始化Utils对象
	U = new(Utils)
}

// DiffTimeInMS 计算当前时间到 start的间隔，以毫秒为单位
func (u *Utils) DiffTimeInMS(start *time.Time) float64 {
	diffNS := time.Since(*start).Nanoseconds()
	return float64(diffNS) / NSToMSUnits
}

// PrintTimeDiffMS 打印start到当前时间的毫秒数
func (u *Utils) PrintTimeDiffMS(start *time.Time) {
	diff := u.DiffTimeInMS(start)
	fmt.Printf("算法耗费时间 %v 毫秒\n", diff)
}

// DiffTimeInNS 计算当前时间到 start的间隔，以纳秒为单位
func (u *Utils) DiffTimeInNS(start *time.Time) float64 {
	diffNS := time.Since(*start).Nanoseconds()
	return float64(diffNS)
}

// PrintTimeDiffNS 打印start到当前时间的纳秒数
func (u *Utils) PrintTimeDiffNS(start *time.Time) {
	diff := u.DiffTimeInNS(start)
	fmt.Printf("算法耗费时间 %v 纳秒\n", diff)
}
