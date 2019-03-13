package vms

import (
	http2 "github.com/aaronize/grandonion/examples/http"
	"github.com/aaronize/grandonion/examples/http/lib"
	"github.com/gin-gonic/gin"
	"kvm/modules/vm-api/models"
	"net/http"
)

func get(c *gin.Context) {
	// 获取参数
	id := c.DefaultQuery("id", "011")
	name := c.DefaultQuery("name", "Jeselin")

	// 数据处理逻辑

	// response
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
		"code": 0,
		"data": "get vms success",
		"id": id,
		"name": name,
	})
}

func create(c *gin.Context) {
	param := &struct {
		UID 	string		`json:"uid"`
	}{}

	if err := c.BindJSON(param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "create vm failed",
			"status": "create_failed",
			"code": lib.CreateVmParamErr,
			"data": "request parameter error",
		})
		return
	}

	// 数据库查询当前状态
	vm := &models.VM{}
	if err := http2.DB.Where("UID = ?", param.UID).Find(vm).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "create vm failed", "status": "create_failed", "code": lib.CreateVmParamErr, "data": ""})
		return
	}
	// 判断当前状态能否被执行。另外，判断必须的配置信息是否齐全


	// Job 约定参数()
	job := http2.Job{Action: CreateVM, UID: param.UID, Parm: []string{}, Success: make(chan string), Fail: make(chan error)}

	// 放入JobQueue队列
	http2.JobQueue <- job
	c.JSON(http.StatusOK, gin.H{"message": "creating vm, please wait...", "status": "creating", "code": lib.EthIsOK, "data": ""})
	return
}

func del(c *gin.Context) {

}

func update(c *gin.Context) {

}

func vmsByPyHostUid(c *gin.Context)  {

}

func vmsByOUid(c *gin.Context) {

}

func detail(c *gin.Context) {

}

func start(c *gin.Context) {

}

func reStart(c *gin.Context) {

}

func shutdown(c *gin.Context) {

}
