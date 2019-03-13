package vms

import (
	"github.com/aaronize/grandonion/examples/http"
	"time"
)

/*

 */

type Engine struct {
	quit 	chan bool
}

func (e *Engine) Start() {
	for {
		select {
		case job := <- http.JobQueue:
			// 拿到job，handle job
			jobHandler(job)

		case <- e.quit:
			// 退出
			return
		}
	}
}

func (e *Engine) Stop() {
	go func() {
		e.quit <- true
	}()
}

func jobHandler(job http.Job) {
	// 设置定时
	t := time.NewTimer(10 * time.Minute)
	// 执行
	executeFunc := job.Action
	// TODO 不用返回error了，因为Job里有了
	go executeFunc(job.Parm)

	select {
	case <- job.Success:
		// 执行成功，生成下一步的job并放入JobQueue中
		nextJob := http.Job{}
		http.JobQueue <- nextJob

		// 释放timer
		t.Stop()
	case fail := <-job.Fail:
		// TODO 执行失败，写库/终止
		fail.Error()

		// 释放timer
		t.Stop()
	case <- job.Finish:
		// 结束

		t.Stop()
	case <- t.C:
		// TODO 超时，写库/终止

	}

}


func engine() {

	for {
		select {
		case job := <- http.JobQueue:
			// 拿到job，handle job
			jobHandler(job)



			//case <- time.After(timeout * time.Minute):
			//	// TODO 这样写超时是不正确的
			//	currentTime := time.Now().Unix()
			//	// 创建超时（从开始创建到最后完成）
			//	vm := &models.VM{}
			//	err := http.DB.Model(vm).Where("UID = ?", currentJob.UID).Updates(
			//		map[string]interface{}{"Info": "执行超时，请检查并尝试重新开始", "LastUpdateTime": currentTime}).Error
			//	if err != nil {
			//		if err == gorm.ErrRecordNotFound {
			//			// 记录日志
			//
			//			return
			//		}
			//	}

		}
	}
}