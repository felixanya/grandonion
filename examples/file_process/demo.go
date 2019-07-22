package main

import (
    "log"
    "os"
)

func main() {
    _, err := os.Stat("config.json")
    if err != nil {
        //log.Println(err)
        if os.IsExist(err) {
            log.Println("File already exist!")
        } else {
            log.Println("File not exist!")
        }
    }
    //fmt.Println("name: ", fi.Name())

}
