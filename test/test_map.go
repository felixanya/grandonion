package main

import (
	"bytes"
	"fmt"
)

func main() {
	var condition bytes.Buffer
	condition.WriteString(" 1=1 ")

	var a = make(map[string]interface{})
	a["OID"] = 1123
	a["PyIP"] = "192.168.168.30"
	for k, v := range a {
		switch v.(type) {
		case int:
			fmt.Println(k, ">>>", v)
			condition.WriteString(fmt.Sprintf(" and %s=%d ", k, v))
		case string:
			fmt.Printf("%s>>>\"%s\"\n", k, v)
			condition.WriteString(fmt.Sprintf(" and %s=\"%s\"", k, v))
		default:
			condition.WriteString("")
		}
	}
	fmt.Println("sql: ", condition.String())
}
