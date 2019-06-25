package main

import (
	"fmt"
	"log"
	"reflect"
)

type TestStruct struct {
	UID 		string 		`json:"uid"`
	OUID 		string 		`json:"ouid"`
	PyUID 		string 		`json:"py_uid"`
	HostName 	string		`json:"host_name"`
	Core 		int 		`json:"core"`
	Mem 		int			`json:"mem"`
	Disk 		int			`json:"disk"`
	OS 			string		`json:"os"`

	InstanceID 	string		`json:"instance_id"`
}

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
	fields := GetFieldName(TestStruct{})
	for i, f := range fields {
		fmt.Printf("第%d个字段名称是：%s\n", i, f)
	}
}
