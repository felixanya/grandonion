package implement2

import "net/http"

func Get(url string, options ...SendOption) (*http.Response, error) {
	return Send("GET", url, options...)
}

func POST(url string, options ...SendOption) (*http.Response, error) {
	return Send("POST", url, options...)
}

func PUT(url string, options ...SendOption) (*http.Response, error) {
	return Send("PUT", url, options...)
}

func DELETE(url string, options ...SendOption) (*http.Response, error) {
	return Send("DELETE", url, options...)
}
