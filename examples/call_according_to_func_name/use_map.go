package main

import (
	"fmt"
)

var FuncNameMapper map[string]func()

func initFuncNameMapper() {
	sw := &Switches{}
	FuncNameMapper = make(map[string]func())
	FuncNameMapper["Login"] = sw.Login
	FuncNameMapper["Change"] = sw.Change
}

type Switches struct {

}

func (s *Switches) Login() {
	fmt.Println("Login")
}

func (s *Switches) Change() {
	fmt.Println("Change")
}

func main() {

	// 初始化map
	initFuncNameMapper()

	// 调用
	FuncNameMapper["Login"]()
	FuncNameMapper["Change"]()


}
