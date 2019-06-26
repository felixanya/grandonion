package model

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type ServerTags []map[string]interface{}

func (st *ServerTags) Scan(src interface{}) error {
	if src == nil {
		return nil
	}
	b, _ := src.([]byte)
	return json.Unmarshal(b, st)
}

func (st ServerTags) Value() (driver.Value, error) {
	return json.Marshal(st)
}

//var _ sql.Scanner = &ServerTags{}
////var _ driver.Value = ServerTags{}

type Server struct {
	gorm.Model
	IP 			string		`gorm:"type:varchar(15)"`
	HostName 	string		`gorm:"type:varchar(40)"`
	ServerTags 	ServerTags	`gorm:"type:json"`
	State 		string		`gorm:"type:varchar(10)"`
	Info 		string		`gorm:"type:varchar(255)"`
}
