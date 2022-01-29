package model

import "net/http"

type ApiResponse struct {
	StatusCode int
	Body       interface{}
}

type ApiResponseWithCookies struct {
	Cookies []*http.Cookie
	ApiResponse
}
