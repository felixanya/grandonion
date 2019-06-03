package main

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func main() {
	uid, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	fmt.Println("v4:", uid.String())
	uid, err = uuid.NewV4()
	fmt.Println("v4:", uid.String())
	uid, err = uuid.NewV4()
	fmt.Println("v4:", uid.String())
	uid, err = uuid.NewV4()
	fmt.Println("v4:", uid.String())
	uid, err = uuid.NewV4()
	fmt.Println("v4:", uid.String())
	fmt.Println("-----------------")
	uid, err = uuid.NewV1()
	fmt.Println("v1:", uid.String())
	uid, err = uuid.NewV1()
	fmt.Println("v1:", uid.String())
	uid, err = uuid.NewV1()
	fmt.Println("v1:", uid.String())
	uid, err = uuid.NewV1()
	fmt.Println("v1:", uid.String())
	uid, err = uuid.NewV1()
	fmt.Println("v1:", uid.String())
	alist := make([]int, 0)
	alist = append(alist, []int{4, 2, 3}...)

	fmt.Println(alist)

}

