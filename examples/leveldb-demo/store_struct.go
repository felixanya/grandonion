package main

import (
	"encoding/json"
	"fmt"
	"github.com/syndtr/goleveldb/leveldb"
	"log"
	"time"
)

type DB struct {
	path 	string
	ldb 	*leveldb.DB
}

type ServerInfo struct {
	Ip 			string
	Sn 			string
	Status 		string
	SubStatus 	string
	CreateAt 	time.Time
	UpdateAt 	time.Time
	PortsInfo	[]PortInfo
	Info 		string
}

type PortInfo struct {
	SwitchManIp 	string
	SwitchName 		string
	PortName 		string
	PortState 		string
	Mac 			string
}

func main() {
	// 目录
	db := NewDB("./leveldb_demo.db")
	if err := db.InitDB(); err != nil {
		log.Fatalf("init leveldb error, %s", err.Error())
	}

	portInfo1 := PortInfo{
		SwitchManIp: "192.168.32.128",
		SwitchName: "SW-N01",
		PortName: "GE1/0/15",
		PortState: "admindown",
		Mac: "52:54:00:0f:bf:e7",
	}
	portInfo2 := PortInfo{
		SwitchManIp: "192.168.32.128",
		SwitchName: "SW-N01",
		PortName: "GE1/0/14",
		PortState: "admindown",
		Mac: "52:54:00:0e:2a:e8",
	}
	portList := []PortInfo{portInfo1, portInfo2}

	serverInfo := &ServerInfo{
		Ip: "192.168.168.11",
		Sn: "SN1234568",
		Status: "isolate",
		SubStatus: "isolating",
		PortsInfo: portList,
		UpdateAt: time.Now(),
		CreateAt: time.Now(),
	}

	if err := db.putServerInfo(serverInfo); err != nil {
		log.Fatal(err)
	}

	ret, err := db.getServerInfo("192.168.168.10")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", ret)

	status, substatus, err := db.getServerStatus("192.168.168.11")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Status: %s, SubStatus: %s\n", status, substatus)

	fmt.Printf("----------------------\n")
	db.Iterator()
}

func NewDB(path string) *DB {
	return &DB{
		path: path,
	}
}

func (d *DB) InitDB() error {
	db, err := leveldb.OpenFile(d.path, nil)
	if err != nil {
		return err
	}
	d.ldb = db

	return nil
}

func (d *DB) putServerInfo(serverInfo *ServerInfo) error {
	bytes, err := json.Marshal(serverInfo)
	if err != nil {
		return err
	}
	return d.ldb.Put([]byte(serverInfo.Ip), bytes, nil)
}

func (d *DB) getServerInfo(ip string) (*ServerInfo, error) {
	bytes, err := d.ldb.Get([]byte(ip), nil)
	if err != nil {
		return nil, err
	}
	serverInfo := &ServerInfo{}
	if err := json.Unmarshal(bytes, serverInfo); err != nil {
		return nil, err
	}
	return serverInfo, nil
}

func (d *DB) getPortList(ip string) ([]PortInfo, error) {
	serverInfo, err := d.getServerInfo(ip)
	if err != nil {
		return nil, err
	}
	return serverInfo.PortsInfo, nil
}

func (d *DB) getServerStatus(ip string) (string, string, error) {
	serverInfo, err := d.getServerInfo(ip)
	if err != nil {
		return "", "", err
	}

	return serverInfo.Status, serverInfo.SubStatus, nil
}

func (d *DB) Iterator() error {
	iter := d.ldb.NewIterator(nil, nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()
		fmt.Printf("key is: %s\n", string(key))

		serverInfo := &ServerInfo{}
		if err := json.Unmarshal(value, serverInfo); err != nil {
			log.Printf("json unmarshal error, %s\n", err.Error())
			continue
		}
		fmt.Printf("Status: %s, SubStatus: %s \n", serverInfo.Status, serverInfo.SubStatus)
	}
	fmt.Printf("end")
	return nil
}
