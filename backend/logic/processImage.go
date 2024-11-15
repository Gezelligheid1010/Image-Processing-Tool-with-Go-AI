package logic

import (
	"backend/dao/api"
	"backend/models"
	"fmt"
	"strings"
)

// UploadImage 上传图像到图床
func UploadImage(image *models.Image) (url string, err error) {
	imageURL, err := api.UploadImageToSMMS(image.OriImage)
	if err != nil {
		fmt.Println("图像上传失败:", err)
		return "", err
	}

	fmt.Println("图像上传成功，URL:", imageURL)

	return imageURL, err
}

// ProcessImage 根据prompt处理图片
func ProcessImage(image *models.Image) (url string, err error) {
	//imageURL, err := api.CallStableDiffusionAPI(image.Prompt, image.OriUrl)
	imageBase64, err := api.InstructPix2Pix(image.Prompt, image.OriUrl)
	if err != nil {
		fmt.Println("图像处理失败:", err)
		return "", err
	}

	fmt.Println("图像处理成功")

	return imageBase64, err
}

// DeleteImage 根据prompt处理图片
func DeleteImage(imageURL string) (err error) {
	imageHash := ExtractHashFromURL(imageURL)
	fmt.Println("imageHash:", imageHash)
	message, err := api.DeleteImageFromSMMS(imageHash, "json")
	if err != nil {
		fmt.Println("删除失败:", err)
		return err
	}
	fmt.Println("删除成功:", message)

	return nil
}

func ExtractHashFromURL(url string) string {
	parts1 := strings.Split(url, "/")
	parts2 := strings.Split(parts1[len(parts1)-1], ".")
	return parts2[0]
}
