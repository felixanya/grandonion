package orders

import "github.com/gin-gonic/gin"

func OrderRouter(or *gin.RouterGroup) {
	//中间件
	or.Use()

	// 获取全部工单
	or.GET("/", Get)
	// 新增一个工单
	or.POST("/", Add)
	// 删除一个工单
	or.DELETE("/:uid", Delete)
	// 更新一个工单
	or.PUT("/:uid", Update)
	//
	or.GET("/:uid/vm-list", GetVms)
}

// 中间件
