package main

import (
	"backend/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	// 添加 CORS 中间件
	router.Use(cors.Default())
	routes.SetupRoutes(router)
	router.Run(":8080") // 启动服务器，监听8080端口
}
