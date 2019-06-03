package main

import (
	"encoding/json"
	"fmt"
)

type Info struct {
	Name 	string	`json:"name"`
	Age 	int		`json:"age"`
	CurrentJob 	string	`json:"current_job"`
}

func main() {

	infos := make([]*Info, 0)

	info := &Info{
		Name: "aaron",
		Age: 20,
		CurrentJob: "engineer",
	}
	info2 := &Info{
		Name: "shawn",
		Age: 10,
		CurrentJob: "student",
	}
	infos = append(infos, info)
	infos = append(infos, info2)

	mp, err := Struct2Map(infos)
	if err != nil {
		panic(err)
	}

	for _, m := range mp {
		delete(m, "age")
	}
	fmt.Println("map: ", mp)
}

func Struct2Map(i interface{}) ([]map[string]interface{}, error) {
	mp := make([]map[string]interface{}, 0)

	bt, err := json.Marshal(i)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(bt, &mp); err != nil {
		return nil, err
	}

	return mp, nil
}

