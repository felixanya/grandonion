package server

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"time"
)

var (
	contentTypeByte = []byte("Content-Type")
	applicationJson = []byte("application/json")
)


func uploadToRedis(ctx *fasthttp.RequestCtx) {
	body := ctx.PostBody()
	res := Resp{
		Meta: Meta{Code: 0, Error: ""},
		Data: nil,
	}

	param := struct {
		Key 		string
		Data 		string
		Expire		time.Duration
	}{}

	if err := json.Unmarshal(body, &param); err != nil {
		// response
		Code = 1
		Error = "parameter error."
		logger.Error(fmt.Sprintf("json unmarshal error."))

		doJSONWrite(ctx, 400, res)
		return
	}

	if err := RC().Set(param.Key, param.Data, param.Expire).Err(); err != nil {
		// response
		Code = 1
		Error = "upload error."
		logger.Error(fmt.Sprintf("upload redis error. %s", err.Error()))

		doJSONWrite(ctx, 400, res)
		return
	}

	// response
	doJSONWrite(ctx, 200, res)
	return
}

func doJSONWrite(ctx *fasthttp.RequestCtx, code int, data interface{}) {
	ctx.Response.Header.SetCanonical(contentTypeByte, applicationJson)
	ctx.Response.SetStatusCode(code)

	if err := json.NewEncoder(ctx).Encode(data); err != nil {
		logger.Error("")
	}
}
