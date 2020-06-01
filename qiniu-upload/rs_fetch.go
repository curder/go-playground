package main

import (
	"fmt"
	"os"

	"github.com/qiniu/api.v7/v7/auth"
	"github.com/qiniu/api.v7/v7/storage"
)

func main() {
	var (
		accessKey = os.Getenv("QINIU_ACCESS_KEY")
		secretKey = os.Getenv("QINIU_SECRET_KEY")
		bucket    = os.Getenv("QINIU_TEST_BUCKET")
	)
	mac := auth.New(accessKey, secretKey)

	cfg := storage.Config{
		// 是否使用https域名进行资源管理
		UseHTTPS: false,
		UseCdnDomains: false,
	}
	// 指定空间所在的区域，如果不指定将自动探测
	// 如果没有特殊需求，默认不需要指定
	//cfg.Zone=&storage.ZoneHuabei
	bucketManager := storage.NewBucketManager(mac, &cfg)

	resURL := os.Getenv("QINIU_TEST_FETCH_URL")
	distName := os.Getenv("QINIU_TEST_SAVE_FILE_NAME")
	// 指定保存的key
	fetchRet, err := bucketManager.Fetch(resURL, bucket, distName)
	if err != nil {
		fmt.Println("fetch error,", err.Error())
	} else {
		fmt.Println(fetchRet.String())
	}
	// 不指定保存的key，默认用文件hash作为文件名
	fetchRet, err = bucketManager.FetchWithoutKey(resURL, bucket)
	if err != nil {
		fmt.Println("fetch error,", err.Error())
	} else {
		fmt.Println(fetchRet.String())
	}
}
