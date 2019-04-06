package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

/*
日志检测log monitoring
 */

const EtcdKey string = "/demodir/demo01_log_mon/192.168.32.128"

type LogConf struct {
	Path 	string	`json:"path"`
	Topic 	string	`json:"topic"`
}

func SetLogConfToEtcd() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"192.168.32.128:2379"},
		DialTimeout: 10 * time.Second,
	})
	if err != nil {
		fmt.Println("新建etcd client 错误，", err.Error())
		return
	}
	defer cli.Close()
	fmt.Println("连接etcd成功...")

	var logConfArr []LogConf

	logConfArr = append(logConfArr, LogConf{
		Path: "/var/log/ipipservice/access.log",
		Topic: "ipipservice_access_log",
	})
	logConfArr = append(logConfArr, LogConf{
		Path: "/var/log/ipipservice/error.log",
		Topic: "ipipservice_error_log",
	})

	data, err := json.Marshal(logConfArr)
	if err != nil {
		fmt.Println("序列化配置信息错误，", err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Put(ctx, EtcdKey, string(data))
	cancel()

	if err != nil {
		fmt.Println("向etcd添加数据失败", err.Error())
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, EtcdKey)
	cancel()
	if err != nil {
		fmt.Println("从etcd获取数据失败", err.Error())
		return
	}

	for _, ev := range resp.Kvs {
		fmt.Printf("%s: %s\n", ev.Key, ev.Value)
		var ret []LogConf
		if err := json.Unmarshal(ev.Value, &ret); err != nil {
			fmt.Printf("获取数据错误, %s\n", err.Error())
			return
		}
		for _, lc := range ret {
			fmt.Printf("获取path：%s, topic: %s\n", lc.Path, lc.Topic)
		}
	}
}

func main() {
	SetLogConfToEtcd()
	fmt.Println("-----------------------------")

	endPoints := []string{"192.168.32.128:2379"}
	cli, err := NewEctdClient(endPoints)
	if err != nil {
		fmt.Printf("创建etcd client失败，%s\n", err.Error())
		return
	}
	defer cli.Close()

	path, data := "/demo/log_mon/config", "demo_value1"
	if err := AnotherExample(path, data, cli); err != nil {
		fmt.Printf("新建数据失败，%s\n", err.Error())
		return
	}
}

func NewEctdClient(endPoints []string) (*clientv3.Client, error) {
	return clientv3.New(clientv3.Config{
		Endpoints: endPoints,
		DialTimeout: 5 * time.Second,
	})
}

func AnotherExample(path, data string, cli *clientv3.Client) error {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	_, err := cli.Put(ctx, path, data)
	cancel()
	if err != nil {
		return err
	}

	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, path)
	cancel()

	for _, ev := range resp.Kvs {
		fmt.Printf("获取数据结果, key: %s, val: %s\n", string(ev.Key), string(ev.Value))
	}

	return nil
}
