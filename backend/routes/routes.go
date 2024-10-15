package routes

import (
	"backend/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/upload", controllers.UploadImage)
		api.POST("/process", controllers.ProcessImage)
	}
}
