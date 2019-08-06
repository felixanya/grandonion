package main

import (
    "log"
    "os"
    "os/exec"
    "runtime"
    "time"
)

func main() {
    runtime.GOMAXPROCS(4)
    defer func() {
        if err := recover(); err != nil {
            log.Println("panic: ", err)
            os.Exit(1)
        }
    }()

    log.Println("daemon start...")
    log.Println("pid: ", os.Getpid())

    watching()
}

func watching() {
    for {
        cmd := exec.Command("/usr/bin/bash", "sleep 60")

        log.Println("start service....")

        if err := cmd.Start(); err != nil {
            log.Println("start service error, ", err.Error())
            time.Sleep(10 * time.Second)
            continue
        }

        _, err := cmd.Process.Wait()
        if err != nil {
            log.Fatalf("wait process error, %s\n", err.Error())
        }
        log.Println("service exited")

        log.Println("service will be restarted in 10 seconds")
        time.Sleep(10 * time.Second)
    }
}
