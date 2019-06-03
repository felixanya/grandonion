package implement2

import (
	"context"
	"io"
	"net/http"
	"time"
)

type sendOptions struct {
	body 			io.Reader
	timeout 		time.Duration
	headers 		map[string]string
	url 			string
	ctx 			context.Context
	retry 			int
}

type SendOption func(*sendOptions)

func SendTimeout(timeout time.Duration) SendOption {
	return func(o *sendOptions) {
		o.timeout = timeout
	}
}

func SendBody(body io.Reader) SendOption {
	return func(o *sendOptions) {
		o.body = body
	}
}

func SendHeaders(headers map[string]string) SendOption {
	return func(o *sendOptions) {
		o.headers = headers
	}
}

func SendRetry(times int) SendOption {
	return func(o *sendOptions) {
		o.retry = times
	}
}

func Send(method, rawurl string, options ...SendOption) (*http.Response, error) {
	// 默认配置
	opts := &sendOptions{
		body: 		nil,
		timeout: 	60 * time.Second,
		headers: 	map[string]string{},
		retry: 		0,
		url: 		"",
		ctx: 		context.Background(),
	}

	// 配置，闭包在调用处执行
	for _, o := range options {
		o(opts)
	}

	client := http.Client{
		Timeout: opts.timeout,
	}

	var resp *http.Response
	var err error
	for {
		resp, err = client.Do(&http.Request{})
		// ...
	}

	return resp, err
}
