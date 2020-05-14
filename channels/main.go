package main

import (
	"fmt"
	"net/url"
)

// "sync"

// 练习
// 启动一个goroutine，生成100个数发送到通道ch1
// 启动一个goroutine，从ch1中取值，计算其值的平方放到通道ch2
// 打印通道ch2中的值

func main() {
	var (
		urlParse *url.URL
		err      error
	)
	u := `/bar`
	if urlParse, err = url.Parse(u); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(urlParse.Path)

}
