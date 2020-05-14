package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/url"
	"os"
	"strings"
)

// 读取文件行，分析出没行的域名，并存入新文件
func main() {
	var (
		path    *string
		file    *os.File
		buf     *bufio.Reader
		line    string
		u       *url.URL
		urls    []string
		domain  string
		newFile *os.File
		err     error
	)

	// 获取文件路径
	path = flag.String("path", "", "请输入文件路径")
	flag.Parse()

	if *path == "" {
		fmt.Println("请输入文件路径")
		return
	}

	// 逐行读取文件
	if file, err = os.Open(*path); err != nil {
		fmt.Printf("Open file err: %s", err.Error())
		return
	}
	defer file.Close() // 关闭文件

	buf = bufio.NewReader(file) // 创建一个缓冲区

	for {
		line, err = buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF { //读取结束，会报EOF
				// 将去重后的结果放到新的文件中
				if newFile, err = os.OpenFile(*path+".domain.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModeAppend); err != nil {
					fmt.Printf("failed to create new file err: %s", err.Error())
				}
				defer newFile.Close()

				for _, domain = range RemoveRepByMap(urls) {
					if _, err = newFile.Write([]byte(domain + "\n")); err != nil {
						fmt.Printf("failed to write file err: %s", err.Error())
						return
					}
				}
				return
			}
			fmt.Printf("Read file err: %s", err)
		}

		// 从字符串中分析出域名
		if u, err = url.Parse(line); err != nil {
			fmt.Printf("parse url is failed, err: %s", err.Error())
		}

		urls = append(urls, u.Hostname())
	}

}

// RemoveRepByMap: 通过map主键唯一的特性过滤重复元素
func RemoveRepByMap(slc []string) []string {
	result := []string{}
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}
