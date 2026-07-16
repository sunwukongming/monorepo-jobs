package services

import (
	"app/config"
	"app/pkg/utils"
	"fmt"
	"sync"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

type OssService struct {
	Endpoint        string
	AccessKeyID     string
	AccessKeySecret string
	Bucket          string
}

var (
	ossOnce    sync.Once
	ossService *OssService
)

func GetOssService() *OssService {
	ossOnce.Do(func() {
		cfg := config.Get().OSS
		ossService = &OssService{
			Endpoint:        cfg.Endpoint,
			AccessKeyID:     cfg.AccessKeyID,
			AccessKeySecret: cfg.AccessKeySecret,
			Bucket:          cfg.Bucket,
		}
	})
	return ossService
}

func (service *OssService) Client() (*oss.Client, error) {
	return oss.New(service.Endpoint, service.AccessKeyID, service.AccessKeySecret)
}

func (service *OssService) PublicURL(objectKey string) string {
	return fmt.Sprintf("https://%s.%s/%s", service.Bucket, service.Endpoint, objectKey)
}

func (service *OssService) PresignUrl(uri string) (string, error) {
	client, err := service.Client()
	if err != nil {
		return "", err
	}
	bucket, err := client.Bucket(service.Bucket)
	if err != nil {
		return "", err
	}
	prefix := fmt.Sprintf("https://%s.%s/", service.Bucket, service.Endpoint)
	if utils.StringStartsWith(uri, prefix) {
		uri = uri[len(prefix):]
	}
	signedURL, err := bucket.SignURL(uri, oss.HTTPGet, 60*60*24)
	if err != nil {
		return "", err
	}
	return signedURL, err
}
