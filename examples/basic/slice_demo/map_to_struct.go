package main

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"
)

type PmInstance struct {
	UUID 	string		`json:"uuid"`
	Owner 	string		`json:"owner"`
	Location 	string	`json:"location"`
	No		int 		`json:"no"`
	Region 	string		`json:"region"`
}

func main() {
	mp := make(map[string]interface{})
	mp["uuid"] = "abcdefgh"
	mp["owner"] = "aaron"
	mp["location"] = "location"
	mp["no"] = 23

	var pi PmInstance
	if err := mapstructure.Decode(mp, &pi); err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("pi: %+v", pi)
}

func Map2Struct(mp map[string]interface{}, st interface{}) (interface{}, error) {
	bt, err := json.Marshal(mp)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(bt, st); err != nil {
		return nil, err
	}

	return st, nil
}
