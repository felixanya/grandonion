package http

import (
	idcs2 "github.com/aaronize/grandonion/examples/servers/http/controller/idcs"
	"github.com/aaronize/grandonion/examples/servers/http/controller/orders"
	pyhosts2 "github.com/aaronize/grandonion/examples/servers/http/controller/pyhosts"
	vms2 "github.com/aaronize/grandonion/examples/servers/http/controller/vms"
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

	orders.OrderRouter(order)
	vms2.VMRoute(vm)
	pyhosts2.PyHostRoute(pyhost)
	idcs2.IdcRoute(idc)
}

// 中间件
