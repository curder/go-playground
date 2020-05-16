package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	var (
		originURL   string
		redirectURL string
		err         error
	)
	originURL = "https://view.vzaar.com/21511487/video"

	if redirectURL, err = getRedirectURL(originURL); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(redirectURL)

}

func getRedirectURL(originURL string) (redirectURL string, err error) {
	var (
		client       *http.Client
		httpRequest  *http.Request
		httpResponse *http.Response
		responseURL  *url.URL
	)
	client = &http.Client{
		CheckRedirect: myCheckRedirect,
	}

	if httpRequest, err = http.NewRequest("GET", originURL, nil); err != nil {
		log.Println(httpRequest.URL)
		return
	}

	if httpResponse, err = client.Do(httpRequest); err != nil {
		log.Println(httpResponse.StatusCode)
	}

	if responseURL, err = httpResponse.Location(); err != nil {
		log.Println(responseURL.String())
		return 
	}

	redirectURL = responseURL.String()

	return
}

func myCheckRedirect(req *http.Request, via []*http.Request) error {
	//自用，将url根据需求进行组合
	if len(via) >= 1 {
		return errors.New("stopped after 1 redirects")
	}
	return nil
}
