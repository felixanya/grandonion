package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	path := os.Args[1]
	handler := func(b []byte) {
		fmt.Println("print line:", string(b))
		return
	}

	if err := ReadFileLine(path, handler); err != nil {
		panic(err)
	}

	//if err := ReadByLine(path, handler); err != nil {
	//	panic(err)
	//}

	return
}

func ReadFileLine(filePath string, handle func([]byte)) error {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err == io.EOF {
				if err := os.Truncate(filePath, 0); err != nil {
					return err
				}
				return nil
			}
			return err
		}
		handle(line)
	}

	return nil
}

func ReadByLine(filePath string, handle func([]byte)) error {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	buf := bufio.NewReader(f)
	for {
		line, err := buf.ReadString('\n')
		if err == nil || err == io.EOF {
			line = strings.TrimSpace(line)
			if line != "" {
				fmt.Println(line)
			}
		}
		if err != nil {
			break
		}
	}
	return nil
}

func ReadFileUseScanner2(filePath string, handle func(*sendMsg)) error {
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

		handle(&sendMsg{
			topic: "",
			value: []byte(line),
		})
	}
	return nil
}
