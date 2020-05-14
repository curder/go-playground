package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "sync"
)

// 对给定的地址资源内容进行对比
func main() {
    var (
        wg       sync.WaitGroup
        suffix   = "app/bower_components/jquery/dist/jquery.min.js"
        url1     = `https://myvip.avatrade.cn/` + suffix
        url2     = `https://myvip.avastocks.cn/` + suffix
        content1 []byte
        content2 []byte
        result1  = make(chan []byte, 1)
        result2  = make(chan []byte, 1)
        err      error
    )
    wg.Add(2)
    go func() {
        defer wg.Done()
        if content1, err = ReadFromRemote(url1); err != nil {
            fmt.Printf("Failed to fetch remote url,err: %s", err.Error())
            return
        }
        result1 <- content1
    }()

    go func() {
        defer wg.Done()
        if content2, err = ReadFromRemote(url2); err != nil {
            fmt.Printf("Failed to fetch remote url,err: %s", err.Error())
            return
        }
        result2 <- content2
    }()

    wg.Wait()

    content1 = <-result1
    content2 = <-result2
    fmt.Println(string(content1) == string(content2))
}

func ReadFromRemote(url string) (content []byte, err error) {
    var (
        resp     *http.Response
        resource io.Reader
    )
    if resp, err = http.Get(url); err != nil {
        return
    }
    resource = resp.Body

    if content, err = ioutil.ReadAll(resource); err != nil {
        fmt.Println("failed to read,err: " + err.Error())
        return
    }

    return
}
