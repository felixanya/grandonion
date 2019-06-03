package ping

import (
	"fmt"
	"testing"
)

func TestPing(t *testing.T) {
	ret := Ping("192.168.168.154", "192.168.152.174", "30.168.152.183")
	fmt.Println("ping result:", ret)
}
