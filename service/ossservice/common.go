package ossservice

import (
	"os"

	"cofmgr/logger"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var bucket *oss.Bucket

func Init() {
	client, err := oss.New(
		os.Getenv("OSS_END_POINT"),
		os.Getenv("OSS_ACCESS_KEY_ID"),
		os.Getenv("OSS_ACCESS_KEY_SECRET"))
	if err != nil {
		logger.Fatal("when get oss client:", err)
	}
	bucket, err = client.Bucket(os.Getenv("OSS_BUCKET"))
	if err != nil {
		logger.Fatal("when create oss bucket:", err)
	}
	if _, err := bucket.IsObjectExist("checkConfig"); err != nil {
		logger.Fatal("please check oss config:", err)
	}
}
