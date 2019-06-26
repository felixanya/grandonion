// +build app2

package server

import "log"

func NewApp2() interface{} {
	log.Println(">>>in NewApp2 return nil>>>")

	return nil
}
