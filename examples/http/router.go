package http

import (
	"github.com/aaronize/grandonion/examples/http/controller/idcs"
	orderRouter "github.com/aaronize/grandonion/examples/http/controller/orders"
	"github.com/aaronize/grandonion/examples/http/controller/pyhosts"
	"github.com/aaronize/grandonion/examples/http/controller/vms"
	"github.com/gin-gonic/gin"
)

func RouteDispatcher(mid *gin.RouterGroup)  {

	// 添加中间件
	mid.Use()

	order := mid.Group("/order")
	vm := mid.Group("/vm")
	//pro := mid.Group("/processor")
	pyhost := mid.Group("/py")
	idc := mid.Group("/idc")

	orderRouter.OrderRouter(order)
	vms.VMRoute(vm)
	pyhosts.PyHostRoute(pyhost)
	idcs.IdcRoute(idc)
}

// 中间件
