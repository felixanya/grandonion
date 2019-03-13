package vms

import "github.com/gin-gonic/gin"

func VMRoute(vr *gin.RouterGroup) {
	// 添加中间件
	vr.Use()

	// 获取列表/分页
	vr.GET("/", get)
	// 新增一条记录
	vr.POST("/", create)

	// 删除一条记录
	vr.DELETE("/:uid", del)
	// 更新一条记录
	vr.PUT("/:uid", update)
	//
	vr.GET("/:uid", detail)
	vr.GET("/:uid/boot", start)
	vr.GET("/:uid/reboot", reStart)
	vr.GET("/:uid/shutdown", shutdown)
	vr.GET("/:uid/")

	//
	vr.GET("/:pyuid/vms", vmsByPyHostUid)
	//
	vr.GET("/:ouid/vms", vmsByOUid)


}
