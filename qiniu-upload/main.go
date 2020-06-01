package main

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/qiniu/api.v7/v7/auth"
	"github.com/qiniu/api.v7/v7/storage"
)

func main() {
	var (
		bucket        = os.Getenv("QINIU_TEST_BUCKET")
		accessKey     = os.Getenv("QINIU_ACCESS_KEY")
		secretKey     = os.Getenv("QINIU_SECRET_KEY")
		mac           *auth.Credentials
		useHTTPS      = false
		useCdnDomains = false
		zone          *storage.Zone
		err           error
		formUploader  *storage.FormUploader

		remoteResourceURL = os.Getenv("QINIU_TEST_FETCH_URL")
		key               = os.Getenv("QINIU_TEST_SAVE_FILE_NAME")
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
