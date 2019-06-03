package main

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	ip := ""
	if err := validateIpFormat(ip); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("ip地址格式合法")
	}
}

func validateIpFormat(target string) error {
	if target == "" {
		return errors.New("ip地址为空")
	}

	ip_port := strings.Split(target, ":")
	if len(ip_port) == 1 {
		return errors.New("格式错误")
	} else if len(ip_port) > 2 {
		return errors.New("格式错误")
	}

	ip := ip_port[0]
	port := ip_port[1]

	if i := net.ParseIP(ip); i == nil {
		//fmt.Println("ip地址格式错误")
		return errors.New("ip地址格式错误")
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		//fmt.Println("prot格式错误")
		return errors.New("port 格式错误")
	}
	if portInt < 0 || portInt > 65535 {
		return errors.New("port不合法，范围0~65535")
	}

	return nil
}
