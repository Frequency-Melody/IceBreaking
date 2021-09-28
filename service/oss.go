package service

import (
	"IceBreaking/config"
	"IceBreaking/log"
	"IceBreaking/response"
	"IceBreaking/response/dto"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"net/url"
	"os"
)

func init() {
	InitOSS()
}

var client *oss.Client
var bucket *oss.Bucket

func InitOSS() {
	// 创建OSSClient实例。
	var err error
	conf := config.Get()
	client, err = oss.New(conf.OSS.EndPoint, conf.OSS.AccessKeyId, conf.OSS.AccessKeySecret)
	if err != nil {
		log.Sugar().Error("阿里云连接失败")
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err = client.Bucket(conf.OSS.BucketName)
	if err != nil {
		log.Sugar().Error("阿里云图库连接失败")
		os.Exit(-1)
	}
}

func UploadFileToOss(filename string, fd io.Reader) response.Response {
	//Signature := base64(hmac-sha1(AccessKeySecret,
	//	VERB + "\n"
	//+ Content-MD5 + "\n"
	//+ Content-Type + "\n"
	//+ Date + "\n"
	//+ CanonicalizedOSSHeaders
	//+ CanonicalizedResource))
	//headers := map[string]string{"Content-Disposition" : "attachment"}
	err := bucket.PutObject(filename, fd)
	baseUrl := "https://bird-hduhelp.oss-cn-hangzhou.aliyuncs.com/"
	pictureUrl := baseUrl + url.QueryEscape(filename)
	if err != nil {
		log.Sugar().Error("文件上传到 OSS 失败，文件名：", filename, ", 错误信息:", err)
		return response.FileUploadToOssFailedError
	}
	return &dto.PictureUrlDto{Url: pictureUrl}
}
