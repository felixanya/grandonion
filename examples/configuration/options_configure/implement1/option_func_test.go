package implement1

import (
	"fmt"
	"testing"
	"time"
)

func TestExecute(t *testing.T) {
	executor, err := Execute(WithEmergency(), WithAction("install"), WithTimeout(30 * time.Second))
	if err != nil {
		panic(err)
	}

	PrintExecutorOptions()
	fmt.Println()
	PrintExecutorInfo()
}
