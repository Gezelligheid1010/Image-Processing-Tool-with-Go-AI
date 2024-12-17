package api

import (
	"backend/settings"
	"bytes"
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
	"log"
	"strings"
	"time"
)

// UploadImageToOSS 上传图片的函数
func UploadImageToOSS(base64Image string) (string, error) {
	//fmt.Println("Endpoint:", settings.Conf.Endpoint)
	//fmt.Printf("Loaded settings: %+v\n", settings.Conf)

	if strings.HasPrefix(base64Image, "data:image") {
		base64Image = strings.Split(base64Image, ",")[1]
	}

	// 解码 Base64 数据
	imageData, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 data: %v", err)
	}

	// 配置 OSS 客户端
	cfg := oss.LoadDefaultConfig().
		WithRegion(settings.Conf.Region).
		//WithEndpoint(settings.Conf.Endpoint).
		WithCredentialsProvider(credentials.NewStaticCredentialsProvider(settings.Conf.AccessKeyID, settings.Conf.AccessKeySecret, ""))
	//WithRegion(settings.Conf.Region)

	client := oss.NewClient(cfg)

	// 生成文件名
	objectName := GenerateUniqueFileName("png")

	// 创建上传对象的请求
	request := &oss.PutObjectRequest{
		Bucket: oss.Ptr(settings.Conf.Bucket), // 存储空间名称
		Key:    oss.Ptr(objectName),           // 对象名称
		Body:   bytes.NewReader(imageData),    // 图片数据
	}

	// 上传文件
	_, err = client.PutObject(context.TODO(), request)
	if err != nil {
		return "", fmt.Errorf("failed to upload image to OSS: %v", err)
	}

	// 生成文件的 URL
	url := fmt.Sprintf("%s/%s", settings.Conf.Endpoint, objectName)

	return url, nil
}

// GenerateUniqueFileName 生成唯一的文件名称
func GenerateUniqueFileName(extension string) string {
	// 获取当前时间戳
	timestamp := time.Now().UnixNano()

	// 生成随机字符串
	randomBytes := make([]byte, 8)
	_, err := rand.Read(randomBytes)
	if err != nil {
		log.Fatalf("failed to generate random bytes: %v", err)
	}
	randomString := hex.EncodeToString(randomBytes)

	// 拼接文件名
	return fmt.Sprintf("%d-%s.%s", timestamp, randomString, extension)
}
