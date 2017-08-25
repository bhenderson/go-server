package utils

import "net/http"

type Response struct {
	Code   int
	Header http.Header
	Body   interface{}

	w http.ResponseWriter
	r *http.Request
}
