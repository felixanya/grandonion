package main

import "log"

func main() {
    log.SetFlags(log.Lshortfile | log.LstdFlags)
    log.SetPrefix("[INFO] ")

    log.Println("testing log")
}

