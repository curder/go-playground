package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

// 过去url地址的页面内容
func main() {
    var (
        u string
    )

    u = "http://127.0.0.1:8080/"
    fmt.Println(string(fromRemote(u)))
    fmt.Println(string(fromFile("/Users/curder/Codes/Go/src/github.com/curder/playground/get-url-contents/main.go")))
}

func fromRemote(rUrl string) (contents []byte, err error) {
    var (
        resp *http.Response
    )
    if resp, err = http.Get(rUrl); err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    if contents, err = ioutil.ReadAll(resp.Body); err != nil {
        return nil, err
    }
    return contents, nil
}

func fromFile(fileName string) (contents []byte, err error) {
    var (
        file *os.File
    )
    if file, err = os.Open(fileName); err != nil {
        return nil, err
    }
    defer file.Close()

    if contents, err = ioutil.ReadAll(file); err != nil {
        return nil, err
    }
    return contents, err
}
