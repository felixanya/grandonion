package main

import (
    "fmt"
    "github.com/peterh/liner"
    "log"
    "os"
    "path/filepath"
    "strings"
)

var (
    history_fn = filepath.Join(os.TempDir(), ".liner_example_history")
    names = []string{"john", "james", "mary", "nancy"}
)

func main() {
    line := liner.NewLiner()
    defer line.Close()

    line.SetCtrlCAborts(true)

    line.SetCompleter(func(line string) (c []string) {
        for _, n := range names {
            if strings.HasPrefix(n, strings.ToLower(line)) {
                c = append(c, n)
            }
        }
        return
    })

    if f, err := os.Open(history_fn); err == nil {
        line.ReadHistory(f)
        f.Close()
    }

    //if name, err := line.Prompt("What's your name?"); err == nil {
    //    log.Println("Got:", name)
    //    line.AppendHistory(name)
    //} else if err == liner.ErrPromptAborted {
    //    log.Println("Aborted")
    //} else {
    //    log.Println("Error reading line:", err)
    //}
    //
    //if f, err := os.Create(history_fn); err != nil {
    //    log.Println("Error writing history file:", err)
    //} else {
    //    line.WriteHistory(f)
    //    f.Close()
    //}

    for {
        if name, err := line.Prompt("cmd> "); err == nil {
            log.Println("Got:", name)
            // execute

            line.AppendHistory(name)

            if name == "exit" {
                fmt.Println("Bye!")
                return
            }
        } else if err == liner.ErrPromptAborted {
            log.Println("Aborted")
            fmt.Println("Bye!")
            return
        } else {
            log.Println("Error reading line:", err)
            return
        }

        if f, err := os.Create(history_fn); err != nil {
            log.Println("Error writing history file:", err)
        } else {
            line.WriteHistory(f)
            f.Close()
        }
    }
}
