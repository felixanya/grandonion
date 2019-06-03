package main

import (
	"encoding/json"
	"fmt"
)

type TestStruct2 struct {
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

func GetFieldNams(target interface{}) ([]string, error) {
	fields := make([]string, 0)

	tsByte, err := json.Marshal(target)
	if err != nil {
		return nil, err
	}
	mp := make(map[string]interface{})

	err = json.Unmarshal(tsByte, &mp)
	if err != nil {
		return nil, err
	}

	for field := range mp {
		fields = append(fields, field)
	}
	return fields, nil
}

func main() {
	ts := TestStruct2{}
	fields, err := GetFieldNams(ts)
	if err != nil {
		panic(err)
	}

	for i, field := range fields {
		fmt.Printf("第%d个字段：%s\n", i, field)
	}
}
