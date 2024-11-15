package controller

import (
	"backend/logic"
	"backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
)

// UploadWorkHandler 上传作品
func UploadWorkHandler(c *gin.Context) {
	// 1、获取参数及校验参数
	var work models.Work
	if err := c.ShouldBindJSON(&work); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "请求数据格式错误"})
		log.Printf("Error binding JSON: %v", err)
		return
	}

	// 2、
	err := logic.UploadWork(&work)
	if err != nil {
		zap.L().Error("上传作品出错", zap.Error(err))
		ResponseError(c, CodeServerBusy)
		fmt.Println("err；", err)
		return
	}
	// 3、返回响应
	ResponseSuccess(c, nil)
}

// DeleteWorksHandler 删除作品
func DeleteWorksHandler(c *gin.Context) {

	// 3、返回响应
	ResponseSuccess(c, nil)
}
