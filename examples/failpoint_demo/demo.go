package main

import (
	"fmt"
	"github.com/pingcap/failpoint"
	"time"
)

func FailPointer() {
	var outerVar = "declare in out scope"
	failpoint.Inject("failpoint-name-for-config", func(val failpoint.Value) {
		fmt.Println("unit-test", val, outerVar)
	})

	time.Sleep(3 * time.Second)
	failpoint.Disable("failpoint-name-for-config")

	time.Sleep(3 * time.Second)
}

func main() {
	FailPointer()
}

