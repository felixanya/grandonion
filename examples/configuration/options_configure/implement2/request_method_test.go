package implement2

import (
	"testing"
	"time"
)

func TestGet(t *testing.T) {
	resp, err := Get("http://localhost:8080/", SendTimeout(30 * time.Second),
		SendHeaders(map[string]string{"Content-Type": "application/json"}))
	// ...
}
