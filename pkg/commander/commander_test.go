package commander

import (
	"log"
	"testing"
)

func TestNewCommandExecutor(t *testing.T) {
	executor := NewCommandExecutor("fping 127.0.0.1")
	if err := executor.Execute(); err != nil {
		log.Printf("execute cmd error, %s. ", err.Error())
		log.Printf("stderr: %s", executor.Stderr())
		log.Printf("stdout: %s", executor.Stdout())
	} else {
		log.Printf("stderr: %s", executor.Stderr())
		log.Printf("stdout: %s", executor.Stdout())
	}
}

func TestNewCommandExecutorWithTimeout(t *testing.T) {
	executor := NewCommandExecutorWithTimeout(
			"fping 39.192.109.23",
			0,
		)
	if err := executor.Execute(); err != nil {
		log.Printf("execute cmd error, %s. ", err.Error())
		log.Printf("stderr: %s", executor.Stderr())
		log.Printf("stdout: %s", executor.Stdout())
	} else {
		log.Printf("stderr: %s", executor.Stderr())
		log.Printf("stdout: %s", executor.Stdout())
	}
}
