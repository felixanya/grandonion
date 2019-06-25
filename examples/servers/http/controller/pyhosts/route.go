package pyhosts

import "github.com/gin-gonic/gin"

func PyHostRoute(pr *gin.RouterGroup) {
	// 中间件
	pr.Use()
}

// 中间件
