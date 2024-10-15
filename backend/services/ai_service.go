package services

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"
)

// UploadImageToSMMS 上传图像到 SM.MS 图床
func UploadImageToSMMS(base64Image, apiToken string) (string, error) {
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
	req.Header.Add("Authorization", apiToken)
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

// CallStableDiffusionAPI 调用Stable Diffusion API处理图像
func CallStableDiffusionAPI(prompt, initImage string, width, height int, steps int, scale float64) (string, error) {
	//url := "https://replicate.com/timothybrooks/instruct-pix2pix/api"
	//url := "https://api.replicate.com/v1/predictions"
	//url := "https://api-inference.huggingface.co/models/timbrooks/instruct-pix2pix"
	//url := "https://stablediffusionapi.com/api/v3/img2img"
	url := "https://modelslab.com/api/v5/controlnet"
	method := "POST"

	// 构建请求数据
	payload := map[string]interface{}{
		"key":              "LfUQGT0IQQrjUUgOGDj8Y2L5Jb3tAKYiRbV2Je33mxIRf5G3heCrPNuEEiz8",
		"controlnet_type":  "canny",
		"controlnet_model": "canny",
		"model_id":         "midjourney",
		"init_image":       initImage,
		"width":            "512",
		"height":           "512",
		"prompt":           prompt,
		"negative_prompt":  nil,
		//"negative_prompt":     "human, unstructure, (black object, white object), colorful background, nsfw",
		"samples":             "1",
		"num_inference_steps": "31",
		"strength":            0.55,
		"guidance_scale":      7.5,
		"scheduler":           "EulerDiscreteScheduler",
	}

	// 序列化请求数据为 JSON
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return "", fmt.Errorf("请求数据序列化失败: %v", err)
	}

	// 创建 HTTP 请求
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonData))
	if err != nil {
		return "", fmt.Errorf("创建请求失败: %v", err)
	}

	// 设置请求头
	req.Header.Add("Content-Type", "application/json")

	// 发起 HTTP 请求
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("请求失败: %v", err)
	}
	defer res.Body.Close()

	// 读取响应数据
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %v", err)
	}

	// 解析 JSON 响应
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("解析响应失败: %v", err)
	}

	// 检查 API 返回的状态
	if result["status"] != "success" {
		return "", fmt.Errorf("AI处理失败: %v", result)
	}

	// 检查是否存在 output 字段并获取图像 URL
	outputArray, ok := result["output"].([]interface{})
	if !ok || len(outputArray) == 0 {
		return "", fmt.Errorf("没有找到生成的图像 URL")
	}

	processedImageURL, ok := outputArray[0].(string)
	if !ok || processedImageURL == "" {
		return "", fmt.Errorf("生成的图像 URL 无效")
	}

	return processedImageURL, nil
}

// FetchGeneratedImage 使用 fetch_result URL 获取生成的图像
func FetchGeneratedImage(fetchResultURL string) (string, error) {
	req, err := http.NewRequest("GET", fetchResultURL, nil)
	if err != nil {
		return "", fmt.Errorf("创建请求失败: %v", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("请求失败: %v", err)
	}
	defer res.Body.Close()

	// 读取响应数据
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("读取响应失败: %v", err)
	}

	// 解析 JSON 响应
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("解析响应失败: %v", err)
	}

	// 检查生成状态和获取图像 URL
	output, ok := result["output"].(string)
	fmt.Println("result:", result)
	if !ok || output == "" {
		return "", fmt.Errorf("生成的图像尚未完成或无法获取")
	}

	return output, nil
}
