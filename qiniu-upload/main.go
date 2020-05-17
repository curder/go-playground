package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/qiniu/api.v7/v7/auth"
	"github.com/qiniu/api.v7/v7/storage"
)

func main() {
	var (
		bucket        = "web-test-private"
		accessKey     = "gz0kbP4_qoYv8lsu1qHg7iWPQJ5Yj4Gq1dgTH63v"
		secretKey     = "DDNztHAYD32hB75OF_vH8S-sQ1k-6xoR04-TcR33"
		mac           *auth.Credentials
		useHTTPS      = true
		useCdnDomains = true
		zone          *storage.Zone
		err           error
		formUploader  *storage.FormUploader

		remoteResourceURL = "https://video.vzaar.com/vzaar/tsj/-V9/target/tsj-V9tTCoiQ_1440_810_2624.mp4?response-content-disposition=inline&X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Credential=AKIAJ74MFWNVAFH6P7FQ%2F20200517%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Date=20200517T082736Z&X-Amz-Expires=3660&X-Amz-SignedHeaders=host&X-Amz-Signature=b676bd41d511f716a5bdd8018d6a669eda9054070a6b33d45527bf83e2bf5239"
		key               = "tsj-V9tTCoiQ_1440_810_2624.mp4"
		// remoteResourceURL = "https://resources.vzaar.com/vzaar/tsj/-V9/target/tsj-V9tTCoiQ_thumb.jpg"
		// key               = "tsj-V9tTCoiQ_thumb.jpg"

		putRet  storage.PutRet
		upToken string
	)

	if zone, err = storage.GetZone(accessKey, bucket); err != nil { // 空间对应的机房
		return
	}

	fmt.Println(zone)

	config := &storage.Config{
		Zone:          zone,          // bucket所在地区
		UseHTTPS:      useHTTPS,      // 是否使用https域名
		UseCdnDomains: useCdnDomains, // 上传是否使用CDN上传加速
	}
	formUploader = storage.NewFormUploader(config)

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac = auth.New(accessKey, secretKey)
	upToken = putPolicy.UploadToken(mac)

	putExtra := storage.PutExtra{}
	u := bytes.Buffer{}
	u.WriteString(remoteResourceURL)
	response, _ := http.Get(u.String())
	defer response.Body.Close()
	readByte, _ := ioutil.ReadAll(response.Body)
	if err = formUploader.Put(context.Background(), &putRet, upToken, key, bytes.NewReader(readByte), int64(len(readByte)), &putExtra); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(putRet.Key)
}
