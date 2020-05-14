package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		timeLayout = `2006-01-02 15:04:05`
		timeString = `2020-05-13T03:27:42+00:00`
		err        error
		result     time.Time
	)

	// 设置当前时区为 Asia/Shanghai
	_, _ = time.LoadLocation("Asia/Shanghai")

	if result, err = time.Parse(time.RFC3339Nano, timeString); err != nil {
		fmt.Printf("parse time err: %s", err.Error())
		return
	}

	fmt.Printf("原始时间：%s\n当前时区时间：%s\n", timeString, result.Local().Format(timeLayout))
}
