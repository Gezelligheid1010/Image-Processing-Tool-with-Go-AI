package api

import (
	"backend/settings"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

// UploadImageToSMMS 上传图像到 SM.MS 图床
func UploadImageToSMMS(base64Image string) (string, error) {
	url := "https://sm.ms/api/v2/upload"

	if strings.HasPrefix(base64Image, "data:image") {
		base64Image = strings.Split(base64Image, ",")[1]
	}

	//fmt.Println("图像:", base64Image)

	// 将 Base64 图片解码为文件字节流
	decodedImage, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return "", fmt.Errorf("Base64 解码失败: %v", err)
	}

	// 创建一个缓冲区和多部分表单写入器
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	// 创建文件字段
	part, err := writer.CreateFormFile("smfile", "image.png")
	if err != nil {
		return "", fmt.Errorf("创建文件字段失败: %v", err)
	}

	_, err = part.Write(decodedImage)
	if err != nil {
		return "", fmt.Errorf("写入文件字段失败: %v", err)
	}

	// 关闭多部分表单写入器
	err = writer.Close()
	if err != nil {
		return "", fmt.Errorf("关闭多部分表单写入器失败: %v", err)
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", fmt.Errorf("创建请求失败: %v", err)
	}

	// 使用 SM.MS 提供的 API Token 进行授权
	req.Header.Add("Authorization", settings.Conf.SmmsToken)
	req.Header.Add("Content-Type", writer.FormDataContentType())

	//fmt.Println("响应内容:", req)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	//fmt.Println("响应内容:", res)
	// 读取响应体
	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %v", err)
	}

	// 解析 JSON 响应
	var result map[string]interface{}
	if err := json.Unmarshal(responseData, &result); err != nil {
		return "", fmt.Errorf("解析响应错误: %v", err)
	}

	// 检查上传状态并获取图像 URL
	success := result["success"].(bool)
	if !success {
		return "", fmt.Errorf("图像上传失败: %v", result["message"])
	}

	data := result["data"].(map[string]interface{})
	imageURL := data["url"].(string)

	return imageURL, nil
}

// DeleteImageFromSMMS 删除 SM.MS 图床上的图像
func DeleteImageFromSMMS(hash string, format string) (string, error) {
	// 构造请求 URL
	url := fmt.Sprintf("https://sm.ms/api/v2/delete/%s?format=%s", hash, format)

	// 创建 HTTP 请求
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("创建请求失败: %v", err)
	}

	// 使用 SM.MS 提供的 API Token 进行授权
	req.Header.Add("Authorization", settings.Conf.SmmsToken)

	// 发送请求
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("发送请求失败: %v", err)
	}
	defer res.Body.Close()

	// 读取响应体
	responseData, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %v", err)
	}

	// 解析 JSON 响应
	var result map[string]interface{}
	if err := json.Unmarshal(responseData, &result); err != nil {
		return "", fmt.Errorf("解析响应错误: %v", err)
	}

	// 检查请求状态并获取响应信息
	success := result["success"].(bool)
	if !success {
		return "", fmt.Errorf("图像删除失败: %v", result["message"])
	}

	return result["message"].(string), nil
}
