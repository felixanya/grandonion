package bridge

import "testing"

func TestNewCircle(t *testing.T) {
	redCircleShape := NewCircle(1, 2, 3, &RedCircle{})
	greenCircleShape := NewCircle(4, 5, 6, &GreenCircle{})

	redCircleShape.Draw()
	greenCircleShape.Draw()

}

/*import (
	"fmt"
	"net/http"
	"testing"
)

func TestClient_Query(t *testing.T) {
	client := &Client{http.DefaultClient}

	cdnReq := &CDNRequest{}
	fmt.Println(client.Query(cdnReq))

	liveReq := &LiveRequest{}
	fmt.Println(client.Query(liveReq))
}
*/