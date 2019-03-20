package model

import "time"

type VM struct {
	UID 		string 		`json:"UID" xorm:"'UID' varchar(40) pk "`
	OUID 		string 		`json:"OUID" xorm:"'OUID' varchar(40)"`
	PyUID 		string 		`json:"PyUID" xorm:"'PyUID' varchar(40)"`
	HostName 	string		`json:"HostName" xorm:"'HostName' varchar(50)"`
	Core 		int 		`json:"Core" xorm:"'Core' int(5)"`
	Mem 		int			`xorm:"'Mem' int(5)"`
	Disk 		int			`xorm:"'Disk' int(5)"`
	Usage 		string		`xorm:"'Usage' varchar(30)"`
	SecUsage 	string		`xorm:"'SecUsage' varchar(30)"`
	OS 			string		`xorm:"'OS' varchar(30)"`
	Owner 		string		`xorm:"'Owner' varchar(20)"`
	Department	string		`xorm:"'Department' varchar(50)"`
	Status 		string		`xorm:"'Status' varchar(20)"`
	SubStatus 	string		`xorm:"'SubStatus' varchar(20)"`
	IDC 		string		`xorm:"'IDC' varchar(20)"`
	Location 	string		`xorm:"'Location' varchar(20)"`
	CreatePerson 	string 	`xorm:"'CreatePerson' varchar(20)"`
	LastUpdatePerson string		`xorm:"'LastUpdatePerson' varchar(20)"`
	CreateTime 		time.Time 		`xorm:"'CreateTime' created"`
	LastUpdateTime 	time.Time 		`xorm:"'LastUpdateTime' updated"`
	InstanceID 	string 		`xorm:"'InstanceID' varchar(40)"`
	//
	//Resource 	Resource 	``
}

