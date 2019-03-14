package order

import "github.com/gin-gonic/gin"


func OrderRouter(r *gin.RouterGroup)  {

	r.GET("/", Get)
	r.POST("/", Add)
	r.DELETE("/", Delete)
	r.PUT("/", Update)
}
