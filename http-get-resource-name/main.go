package main

import (
	"fmt"
	"net/http"
	"path"
)

// 通过地址栏获取对应的文件名
func main() {
	var (
		originUrl string
	)
	originUrl = "https://resources.vzaar.com/vzaar/tmS/hDx/target/tmShDxWnGQT8_thumb.jpg"
	r, _ := http.NewRequest("GET", originUrl, nil)

	fmt.Println(path.Base(r.URL.Path))
}
