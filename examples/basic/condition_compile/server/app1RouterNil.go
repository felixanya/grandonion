// +build !app1

package server

import "log"

func NewApp1() interface{} {
	log.Println(">>>in NewApp1 return nil>>>")
	return nil
}
