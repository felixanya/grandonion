package model

import (
	"database/sql/driver"
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type VIPerson struct {
	gorm.Model
	RealName 	string		`gorm:"type:varchar(30)"`
	NicName 	string		`gorm:"type:varchar(30)"`
	Age 		int			`gorm:"type:int(3)"`

	OtherInfo 	*InfoMap 	`gorm:"type:json"`
}

type InfoMap struct {
	Src 	map[string]string
	Valid 	bool
}

func (im *InfoMap) Scan(val interface{}) error {
	if val == nil {
		im.Src, im.Valid = make(map[string]string), false
		return nil
	}
	t := make(map[string]string)
	if err := json.Unmarshal(val.([]byte), &t); err != nil {
		return err
	}
	im.Valid = true
	im.Src = t
	return nil
}

func (im *InfoMap) Value() (driver.Value, error) {
	if im == nil {
		return nil, nil
	}
	if !im.Valid {
		return nil, nil
	}
	dv, err := json.Marshal(im.Src)
	return dv, err
}