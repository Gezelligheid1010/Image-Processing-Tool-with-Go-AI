package utils

import (
	"encoding/base64"
	"io/ioutil"
	"net/http"
)

// ConverUrl2Base64 从URL加载图片数据并将其转换为Base64格式
func ConverUrl2Base64(url string) (string, error) {
	// 发起HTTP请求
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 检查HTTP状态码是否为200（成功）
	if resp.StatusCode != http.StatusOK {
		return "", err
	}

	// 读取图片的二进制数据
	imageData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// 将图片数据转换为Base64编码
	base64Data := base64.StdEncoding.EncodeToString(imageData)

	// 可选：添加数据类型前缀，方便在HTML中显示
	mimeType := http.DetectContentType(imageData)
	base64String := "data:" + mimeType + ";base64," + base64Data

	return base64String, nil
}
