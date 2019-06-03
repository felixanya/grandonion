package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fn, err := os.Open("examples/file_process/pingmesh_data.txt")
	if err != nil {
		panic(err)
	}
	defer fn.Close()

	buf := bufio.NewReader(fn)
	for {
		line, _, err := buf.ReadLine()
		fmt.Println("line:", string(line))
		if err != nil {
			if err == io.EOF {
				fmt.Println("end")
				return
			}
			panic(err)
		}

	}

}
