package json_data

import "time"

type TableXorm struct {
	ID 		int64 		`json:"id" xorm:"'id' autoincr"`
	UID 	string		`json:"uid" xorm:"'uid' pk"`
	CreateTime 	time.Time 	`json:"create_time" xorm:"'create_time' created"`
	UpdateTime 	time.Time 	`json:"update_time" xorm:"'update_time' updated"`
	Info 	string		`json:"info" xorm:"'info'"`
}

