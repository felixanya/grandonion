package http

import (
	"github.com/aaronize/grandonion/examples/jobque/http/controller/order"
	"github.com/aaronize/grandonion/examples/jobque/http/controller/processor"
	"github.com/gin-gonic/gin"
)

func orderRouter(r *gin.RouterGroup)  {

	r.GET("/", order.Get)
	r.POST("/", order.Add)
	r.DELETE("/", order.Delete)
	r.PUT("/", order.Update)
}

func vmRouter(r *gin.RouterGroup) {

}

func resourceRouter(r *gin.RouterGroup) {

}

func processorRouter(r *gin.RouterGroup) {
	r.POST("/msg", processor.Processor)
}
