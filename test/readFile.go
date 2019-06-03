package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fileName := "./output.log"
	go func(name string) {
		fileObj, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			fmt.Println("Failed to open the file", err.Error())
			os.Exit(2)
		}
		defer fileObj.Close()

		for i := 0; i < 100; i++ {
			if _, err := fileObj.WriteString(fmt.Sprintf("line: %d", i)); err != nil {
				fmt.Println("Failed to writing to the file with os.OpenFile function")
			}
		}
	}(fileName)
	//
	time.Sleep(20 * time.Second)
}
