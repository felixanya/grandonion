package main

import (
	"fmt"
	"strings"
)

func main() {
	uid := "123abcdefg"
	cmd := fmt.Sprintf(`virsh dominfo %s | grep -w "Name" | awk '{print $2}' | while read name;
    do
        ps aux | grep ${name} | grep -v grep|grep qemu-kvm|awk '{print $2}' | while read pid;
        do
            netstat -tulnp|grep ${pid}/qemu-kvm|awk '{print $4}'|awk -F ':' '{print $2}' | while read port;
            do
                echo ${port};
            done;
        done;
    done
	`, uid)
	fmt.Println(cmd)
	fmt.Println("--------------------")

	tar := `Id:             103
Name:           192-168-168-154
UUID:           a432ca90-b1a0-11e8-8c4e-5270000bad43
OS Type:        hvm
State:          running
CPU(s):         2
CPU time:       4805857.3s
Max memory:     4096000 KiB
Used memory:    4096000 KiB
Persistent:     yes
Autostart:      disable
Managed save:   no
Security model: none
Security DOI:   0`
	ret := process_multi(tar)
	fmt.Println("result: ", ret)
}

func process_multi(tar string) map[string]string {
	ret := make(map[string]string)
	tarList := strings.Split(tar, "\n")
	for _, line := range tarList {
		if strings.Contains(line, "Name") {
			val := strings.Trim(strings.Split(line, ":")[1], " ")
			ret["HostName"] = val
		} else if strings.Contains(line, "UUID") {
			val := strings.Trim(strings.Split(line, ":")[1], " ")
			ret["UID"] = val
		} else if strings.Contains(line, "State") {
			val := strings.Trim(strings.Split(line, ":")[1], " ")
			ret["State"] = val
		} else if strings.Contains(line, "CPU") {
			val := strings.Trim(strings.Split(line, ":")[1], " ")
			ret["Core"] = val
		} else if strings.Contains(line, "memory") {
			val := strings.Trim(strings.Split(line, ":")[1], " ")
			ret["Mem"] = val
		}
	}
	return ret
}
