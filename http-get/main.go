package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main()  {
	var (
		resp *http.Response
		result []byte
		err error
	)
	
	// Request and check error
	if resp, err = http.Get("https://view.vzaar.com/21563208/video"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	defer resp.Body.Close()

	// Read Body and check error
    if result, err = ioutil.ReadAll(resp.Body); err != nil {
        fmt.Println(err)
        os.Exit(1)
	}
	
    // Print response html : conver byte to string
    fmt.Println(string(result))
}