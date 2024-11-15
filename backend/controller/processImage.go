package controller

import (
	"backend/logic"
	"backend/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
)

// ProcessImageHandler 调用AI接口进行图像处理
func ProcessImageHandler(c *gin.Context) {
	// 1、获取参数及校验参数
	var image models.Image
	if err := c.ShouldBindJSON(&image); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求数据格式错误"})
		log.Printf("Error binding JSON: %v", err)
		return
	}

	// 获取作者ID，当前请求的UserID(从c取到当前发请求的用户ID)
	userID, err := getCurrentUserID(c)
	if err != nil {
		zap.L().Error("GetCurrentUserID() failed", zap.Error(err))
		ResponseError(c, CodeNotLogin)
		return
	}
	image.AuthorId = userID
	// 2、上传图像
	origenUrl, err := logic.UploadImage(&image)
	if err != nil {
		zap.L().Error("上传图像出错", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}
	image.OriUrl = origenUrl

	//image.OriUrl = "https://s2.loli.net/2024/11/07/6hImzaY4GWTqP7e.png"

	// 3、处理图像
	processedImage, err := logic.ProcessImage(&image)
	if err != nil {
		zap.L().Error("AI处理图片出错", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		return
	}

	//// 4、删除原图
	//err = logic.DeleteImage(image.OriUrl)
	//if err != nil {
	//	zap.L().Error("删除原图出错", zap.Error(err))
	//	ResponseError(c, CodeServerBusy)
	//	return
	//}
	// 5、返回响应
	ResponseSuccess(c, processedImage)
}
