package models

/*
	工单表
 */
type Order struct {
	UID 		string		`gorm:"column:UID;primary_key;type:varchar(40)"`
	IDC 		string		`gorm:"column:IDC;type:varchar(20)"`
	Location 	string		`gorm:"column:Location;type:varchar(20)"`
	Num 		int			`gorm:"column:Num;type:int(5)"`
	Core 		int			`gorm:"column:Core;type:int(5)"`
	Mem 		int			`gorm:"column:Mem;type:int(5)"`
	Disk 		int			`gorm:"column:Disk;type:int(5)"`
	OS 			string		`gorm:"column:OS;type:varchar(20)"`
	Owner 		string		`gorm:"column:Owner;type:varchar(20)"`
	Department 	string		`gorm:"column:Department;type:varchar(50)"`
	Usage 		string		`gorm:"column:Usage;type:varchar(30)"`
	SecUsage	string		`gorm:"column:SecUsage;type:varchar(30)"`
	Origin 		string		`gorm:"column:Origin;type:varchar(20)"`
	CreateTime 	int			`gorm:"column:CreateTime;type:int(10)"`
	TaskID 		string		`gorm:"column:TaskID;type:varchar(35)"`
	ItemID 		string		`gorm:"column:ItemID;type:varchar(35)"`
	Operator 	string		`gorm:"column:Operator;type:varchar(20)"`
	Status 		string		`gorm:"column:Status;type:varchar(20)"`
	Set 		int			`gorm:"column:Set;type:int(1)"`
	InstanceID 	string		`gorm:"column:InstanceID;type:varchar(35)"`

	Vms 		[]VM		`gorm:"foreignkey:OUID;association_foreignkey:UID"`
}

/*
	虚机表
 */
type VM struct {
	UID 		string		`gorm:"column:UID;primary_key;type:varchar(40)"`
	OUID 		string		`gorm:"column:OUID;type:varchar(40)"`
	PyUID 		string		`gorm:"column:PyUID;type:varchar(40)"`
	HostName 	string		`gorm:"column:HostName;type:varchar(50)"`
	Core 		int			`gorm:"column:Core;type:int(5)"`
	Mem 		int			`gorm:"column:Mem;type:int(5)"`
	Disk 		int			`gorm:"column:Disk;type:int(5)"`
	Usage 		string		`gorm:"column:Usage;type:varchar(30)"`
	SecUsage 	string		`gorm:"column:SecUsage;type:varchar(30)"`
	OS 			string		`gorm:"column:OS;type:varchar(20)"`
	Owner 		string		`gorm:"column:Owner;type:varchar(20)"`
	Department	string		`gorm:"column:Department;type:varchar(50)"`
	Status 		string		`gorm:"column:Status;type:varchar(20)"`
	SubStatus 	string		`gorm:"column:SubStatus;type:varchar(20)"`
	IDC 		string		`gorm:"column:IDC;type:varchar(20)"`
	Location 	string		`gorm:"column:Location;type:varchar(20)"`
	LastUpdatePerson 	string		`gorm:"column:LastUpdatePerson;type:varchar(20)"`
	LastUpdateTime 		int			`gorm:"column:LastUpdateTime;type:int(10)"`
	CreatePerson		string		`gorm:"column:CreatePerson;type:varchar(20)"`
	CreateTime 	int			`gorm:"column:CreateTime;type:int(10)"`
	Info 		string 		`gorm:"column:Info;type:varchar(255)"`

	Resource 	Resource	`gorm:"foreignkey:UID;association_foreignkey:UID"`
}

/*
	资源表
 */
type Resource struct {
	UID 	string		`gorm:"column:UID;primary_key;type:varchar(35)"`
	VmUID 	string		`gorm:"column:VmUID;type:varchar(35)"`
	PyIP 	string		`gorm:"column:PyIP;type:varchar(20)"`
	Rack 	string		`gorm:"column:Rack;type:varchar(20)"`
	Mac 	string		`gorm:"column:Mac;type:varchar(60)"`
	IP 		string		`gorm:"column:IP;type:varchar(20)"`
	Status 	int			`gorm:"column:Status;type:int(1)"`
}

/*
	物理机表
 */
type PyHost struct {
	UID 		string		`gorm:"column:UID;primary_key;type:varchar(40)"`
	SN 			string		`gorm:"column:SN;type:varchar(35)"`
	IDC 		string		`gorm:"column:IDC;type:varchar(20)"`

	HostName 	string		`gorm:"column:HostName;type:varchar(35)"`
	MachineType string		`gorm:"column:MachineType;type:varchar(50)"`
	Core 		int			`gorm:"column:Core;type:int(10)"`
	UsedCore	int			`gorm:"column:UsedCore;type:int(10)"`
	FreeCore 	int			`gorm:"column:FreeCore;type:int(10)"`
	Mem 		int			`gorm:"column:Mem;type:int(10)"`
	UsedMem 	int			`gorm:"column:UsedMem;type:int(10)"`
	FreeMem		int			`gorm:"column:FreeMem;type:int(10)"`
	VmCore 		int			`gorm:"column:VmCore;type:int(10)"`
	VmMem 		int			`gorm:"column:VmMem;type:int(10)"`
	Location    string		`gorm:"column:Location;type:varchar(35)"`
	Rack 		string		`gorm:"column:Rack;type:varchar(35)"`
	Locked 		int			`gorm:"column:Locked;type:int(1)"`
	Remark 		string		`gorm:"column:Remark;type:varchar(255)"`
	UsageType 	string		`gorm:"column:UsageType;type:varchar(35)"`
	Status  	int			`gorm:"column:Status;type:int(1)"`
	ReportTime 	int			`gorm:"column:ReportTime;type:int(10)"`

	Vms 		[]VM		`gorm:"foreignkey:PyUID;association_foreignkey:UID"`
}

/*
	逻辑机房表
 */
type IdcConf struct {
	IDC 		string		`gorm:"column:IDC;primary_key;type:varchar(20)"`
	IpSeg 		string		`gorm:"column:IpSeg;type:varchar(100)"`
	Route 		string		`gorm:"column:Route;type:varchar(100)"`
	MacSeg 		string		`gorm:"column:MacSeg;type:varchar(100)"`
	Eth0 		string		`gorm:"column:Eth0;type:varchar(20)"`
	Eth1 		string		`gorm:"column:Eth1;type:varchar(20)"`
	Eth2 		string		`gorm:"column:Eth2;type:varchar(20)"`
	Location 	string		`gorm:"column:Location;type:varchar(50)"`
	Set 		int			`gorm:"column:Set;type:int(1)"`
	Info 		string		`gorm:"column:Info;type:varchar(255)"`

	PyHosts 	[]PyHost 	`gorm:"foreignkey:IDC;association_foreignkey:IDC"`
}
