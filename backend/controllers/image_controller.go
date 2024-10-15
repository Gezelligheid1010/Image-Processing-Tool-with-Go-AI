package controllers

import (
	"backend/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// UploadImage 上传图像处理
func UploadImage(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "图像上传失败"})
		return
	}
	// 将文件保存到服务器
	filePath := "./uploads/" + file.Filename
	c.SaveUploadedFile(file, filePath)

	c.JSON(http.StatusOK, gin.H{"filePath": filePath})
}

// ProcessImage 调用AI接口进行图像处理
func ProcessImage(c *gin.Context) {
	var json struct {
		Prompt    string  `json:"prompt"`
		InitImage string  `json:"init_image"`
		Width     string  `json:"width"`
		Height    string  `json:"height"`
		Steps     string  `json:"num_inference_steps"`
		Scale     float64 `json:"guidance_scale"`
	}
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求数据格式错误"})
		log.Printf("Error binding JSON: %v", err)
		return
	}

	width, _ := strconv.Atoi(json.Width)
	height, _ := strconv.Atoi(json.Height)
	steps, _ := strconv.Atoi(json.Steps)

	apiToken := "4LrqBiEgBtJDaTZSbfmk56xRHd1kYR7r"

	imageURL, err := services.UploadImageToSMMS(json.InitImage, apiToken)
	if err != nil {
		fmt.Println("图像上传失败:", err)
	} else {
		fmt.Println("图像上传成功，URL:", imageURL)
	}

	//processedImage, err := services.CallStableDiffusionAPI(json.Prompt, imageURL, width, height, steps, json.Scale)
	//if err != nil {
	//	log.Printf("Error processing image: %v", err)
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "AI处理失败"})
	//	return
	//}

	// 第一步：发送初始请求
	processedImageURL, err := services.CallStableDiffusionAPI(json.Prompt, imageURL, width, height, steps, json.Scale)
	if err != nil {
		fmt.Println("Error:", err)
		return
	} else {
		fmt.Println("Generated Image URL:", processedImageURL)
	}

	//fmt.Println("Processing image, please wait...")
	//
	//// 等待 ETA 时间，然后请求获取生成的图像
	//time.Sleep(150 * time.Second) // 等待时间可以根据返回的 eta 动态调整
	//
	//// 第二步：使用 fetch_result URL 获取图像
	//processedImageURL, err := services.FetchGeneratedImage(fetchResultURL)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//} else {
	//	fmt.Println("Generated Image URL:", processedImageURL)
	//}
	//fmt.Println("Generated Image URL:", processedImageURL)

	c.JSON(http.StatusOK, gin.H{"processedImage": processedImageURL})
}
