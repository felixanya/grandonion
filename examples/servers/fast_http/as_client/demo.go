package as_client

import (
	"fmt"
	"github.com/valyala/fasthttp"
)

func main() {
	url := "http://httpbin.org/post?key=123"

	req := fasthttp.AcquireRequest()
	resp := fasthttp.AcquireResponse()

	defer func() {
		fasthttp.ReleaseResponse(resp)
		fasthttp.ReleaseRequest(req)
	}()

	req.Header.SetContentType("application/json")
	req.Header.SetMethod("POST")

	req.SetRequestURI(url)

	if err := fasthttp.Do(req, resp); err != nil {
		fmt.Println("request failed, ", err.Error())
		return
	}

	b := resp.Body()
	fmt.Printf("result: \r\n%s", string(b))
}
