package bridge

import "fmt"

type DrawAPI interface {
	DrawCircle(int, int, int)
}

// red circle
type RedCircle struct {
}

func (rc *RedCircle) DrawCircle(radius, x, y int)  {
	fmt.Printf("Drawing circle[color: red, radius: %d, x: %d, y: %d]\n", radius, x, y)
}

// green circle
type GreenCircle struct {
	Color 	string
}

func (gc *GreenCircle) DrawCircle(radius, x, y int) {
	fmt.Printf("Drawing circle[color: green, radius: %d, x: %d, y: %d]\n", radius, x, y)
}

// shape
type IShape interface {
	Draw()
}

type Circle struct {
	radius int
	x int
	y int
	DrawAPI DrawAPI
}

func (c *Circle) Draw() {
	c.DrawAPI.DrawCircle(c.radius, c.x, c.y)
}

func NewCircle(radius, x, y int, drawAPI DrawAPI) IShape {
	return &Circle{
		radius: radius,
		x: x,
		y: y,
		DrawAPI: drawAPI,
	}
}




/*import "net/http"

type Request interface {
	HttpRequest() (*http.Request, error)
}

//
type Client struct {
	Client *http.Client
}

func (c *Client) Query(req Request) (resp *http.Response, err error) {
	httpReq, _ := req.HttpRequest()
	resp, err = c.Client.Do(httpReq)
	return
}

//
type CDNRequest struct {

}

func (cdn *CDNRequest) HttpRequest() (*http.Request, error) {
	return http.NewRequest("GET", "/cdn", nil)
}

//
type LiveRequest struct {

}

func (lr *LiveRequest) HttpRequest() (*http.Request, error) {
	return http.NewRequest("GET", "/live", nil)
}*/