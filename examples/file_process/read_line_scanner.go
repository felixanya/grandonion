package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	path := os.Args[1]
	handler := func(b []byte) {
		fmt.Println("print line:", string(b))
		return
	}

	if err := ReadFileUseScanners(path, handler); err != nil {
		panic(err)
	}

	return
}

func ReadFileUseScanners(filePath string, handle func([]byte)) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		handle([]byte(line))
	}
	//if err := os.Truncate(filePath, 0); err != nil {
	//	return err
	//}
	return nil
}
