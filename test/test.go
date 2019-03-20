package main

import (
	"fmt"
	"github.com/satori/go.uuid"
	"strings"
	"time"
)

func main2() {

	timer := time.NewTimer(3 * time.Second)
	quit := make(chan bool)

	go func() {
		fmt.Println("start execute goroutine")

		select {
		case <-timer.C:
			fmt.Println("Timer has expired.")
		case <-quit:
			fmt.Println("quit")
			return
		default:
			fmt.Println("execute timer.Stop")
			timer.Stop()
		}

	}()

	//timer.Stop()
	time.Sleep(60 * time.Second)
	//select {}
}

func main() {
	keyword := "192.168.152.174"
	sql := fmt.Sprintf(`and concat(ifnull(Mac,""),ifnull(Department,""),ifnull(Usages,""),ifnull(SecUsages,""),ifnull(Status,""),ifnull(IDC,""),ifnull(IP,""),ifnull(HostName,""),ifnull(Mem,""),ifnull(PyHostIP,""),ifnull(Owner,""),ifnull(Disk,""),ifnull(OS,""),ifnull(Remark,""),ifnull(CreateTime,"")) like "%%%s%%"`, keyword)
	fmt.Println("slq is:", sql)

	orderby := "-Core"
	if strings.HasPrefix(orderby, "-") {
		fmt.Println(orderby, "has prefix '-'")
		orderby = strings.Split(orderby, "-")[1]
		fmt.Println("order by", orderby)
	}

	id, _ := uuid.NewV4()

	fmt.Println("generate uuid:", id)
}