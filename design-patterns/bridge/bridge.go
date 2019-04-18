package bridge

import "fmt"

/*
强关联改换成弱关联，将两个角色之间的继承关系改为关联关系。由继承（或颜色在每个类中分别定义）改为组合。
以形状、颜色等属性分为多个维度，每个维度一个类型，这些类型通过组合的方式关联在一起。复用的方式减少代码量。
 */
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