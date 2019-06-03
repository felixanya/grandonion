package main

import (
	"fmt"
	"reflect"
)

type ControllerMapper map[string]reflect.Value

var crtlMapper ControllerMapper

type Routers struct {

}

func (r *Routers) Login(msg string) {
	fmt.Println("Login:", msg)
}

func (r *Routers) ChangeName(msg string) {
	fmt.Println("ChangeName:", msg)
	msg = msg + " Changed"
}

func main() {
	var rt Routers

	crtlMapper = make(ControllerMapper)

	vf := reflect.ValueOf(&rt)
	vfType := vf.Type()

	mNum := vf.NumMethod()
	fmt.Println("NumMethod:", mNum)

	for i := 0; i < mNum; i++ {
		mName := vfType.Method(i).Name
		fmt.Println("index:", i, "MethodName:", mName)
		crtlMapper[mName] = vf.Method(i)
	}

	crtlMapper["Login"].Call([]reflect.Value{reflect.ValueOf("hello Go!")})

	crtlMapper["ChangeName"].Call([]reflect.Value{reflect.ValueOf("Hello Go2!")})
}
