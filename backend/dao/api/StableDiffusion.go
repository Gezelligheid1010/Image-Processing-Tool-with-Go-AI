package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// CallStableDiffusionAPI 调用Stable Diffusion API处理图像
func CallStableDiffusionAPI(prompt, initImage string) (string, error) {
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

	//fmt.Println("body:", body)
	// 解析 JSON 响应
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", fmt.Errorf("解析响应失败: %v", err)
	}

	fmt.Println("result:", result)
	// 检查并获取图像 URL
	var imageURL string
	if futureLinks, ok := result["future_links"].([]interface{}); ok && len(futureLinks) > 0 {
		// 使用 future_links 中的第一个链接
		imageURL, _ = futureLinks[0].(string)
	} else if output, ok := result["output"].([]interface{}); ok && len(output) > 0 {
		// 使用 output 中的第一个链接
		imageURL, _ = output[0].(string)
	} else {
		return "", fmt.Errorf("没有找到图像 URL")
	}

	// 确保获得有效的 imageURL
	if imageURL == "" {
		return "", fmt.Errorf("图像 URL 无效")
	}

	// 轮询图像 URL，直到图片就绪
	maxRetries := 20
	for i := 0; i < maxRetries; i++ {
		resp, err := http.Get(imageURL)
		if err == nil && resp.StatusCode == 200 {
			resp.Body.Close()
			return imageURL, nil
		}
		resp.Body.Close()
		time.Sleep(5 * time.Second) // 每次请求后等待 5 秒
	}

	return "", fmt.Errorf("图像生成超时，请稍后再试")
}
