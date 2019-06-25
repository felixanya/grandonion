package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

/*
使用map存取json数据，很灵活可以随时增减字段
*/
type Attribute []map[string]interface{}

func (a *Attribute) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	b, _ := src.([]byte)
	return json.Unmarshal(b, a)
}

func (a Attribute) Value() (driver.Value, error) {
	return json.Marshal(a)
}

//var _ sql.Scanner = &Attribute{}
//var _ driver.Value = Attribute{}

type MicroBlog struct {
	ID 			int
	Author 		string
	Content 	string
	Attribute 	Attribute
	CreateAt 	*time.Time
	UpdateAt 	*time.Time
}
