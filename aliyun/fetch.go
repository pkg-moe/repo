package aliyun

import (
	"encoding/json"
	"io"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type AliyunInfo struct {
	accessKey string // aliyun access key
	secretKey string // aliyun secret key
	endpoint  string // aliyun oss endpoint
	bucket    string // aliyun oss bucket
	dir       string // aliyun oss dir
}

func NewAliyunInfo(accessKey, secretKey, endpoint, bucket, dir string) *AliyunInfo {
	return &AliyunInfo{
		accessKey: accessKey,
		secretKey: secretKey,
		endpoint:  endpoint,
		bucket:    bucket,
		dir:       dir,
	}
}

func (a *AliyunInfo) Fetch(fileName string) ([]byte, error) {
	client, err := oss.New(a.endpoint, a.accessKey, a.secretKey)
	if err != nil {
		return nil, err
	}

	bucket, err := client.Bucket(a.bucket)
	if err != nil {
		return nil, err
	}

	body, err := bucket.GetObject(a.dir + fileName)
	if err != nil {
		return nil, err
	}
	defer body.Close()

	data, err := io.ReadAll(body)
	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, ErrInvalidJSON
	}

	if !json.Valid(data) {
		return nil, ErrInvalidJSON
	}

	return data, nil
}
