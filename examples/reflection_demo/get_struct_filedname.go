package main

import (
	"fmt"
	"github.com/aaronize/grandonion/examples/reflection_demo/model"
	"log"
	"reflect"
)

/*
使用反射获取结构体字段名称
 */
func GetFieldName(structName interface{}) []string {
	t := reflect.TypeOf(structName)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		log.Println("input is not struct")
		return nil
	}

	fieldNum := t.NumField()
	result := make([]string, 0, fieldNum)
	for i := 0; i < fieldNum; i++ {
		result = append(result, t.Field(i).Name)
	}
	return result
}


func main() {
	fields := GetFieldName(model.VM{})
	for i, f := range fields {
		fmt.Printf("第%d个字段名称是：%s\n", i, f)
	}
}
