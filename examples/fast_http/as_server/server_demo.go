package as_server

import (
	"encoding/json"
	"fmt"
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)

func Index(ctx *fasthttp.RequestCtx) {
	fmt.Fprint(ctx, "Welcome!\n")
}

func Hello(ctx *fasthttp.RequestCtx) {
	fmt.Fprintf(ctx, "hello, %s!\n", ctx.UserValue("name"))
}

func Process(ctx *fasthttp.RequestCtx) {
	body := ctx.PostBody()
	fmt.Fprintf(ctx, "Body: %s", body)

	param := struct {
		ID 		int 	`json:"id"`
		Action 	string	`json:"action"`
		Data 	[]string `json:"data"`
	}{}
	json.Unmarshal(body, &param)
	fmt.Printf("%v\n", param)
	fmt.Printf("---------------------\n")
	fmt.Println("ID: ", param.ID)
	fmt.Println("Action: ", param.Action)
	fmt.Println("Data: ", param.Data)
}

func main() {
	router := fasthttprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.POST("/", Process)

	if err := fasthttp.ListenAndServe("0.0.0.0:8090", router.Handler); err != nil {
		fmt.Println("start fasthttp failed, ", err.Error())
	}
}
