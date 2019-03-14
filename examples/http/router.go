package http

import (
	"github.com/aaronize/grandonion/examples/http/controller/processor"
	"github.com/gin-gonic/gin"
)



func vmRouter(r *gin.RouterGroup) {

}

func resourceRouter(r *gin.RouterGroup) {

}

func processorRouter(r *gin.RouterGroup) {
	r.POST("/msg", processor.Processor)
}
