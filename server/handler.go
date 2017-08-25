package server

import "net/http"

type Handler func(*Response)
type middleware func(Handler) Handler
type middlewareStack []middleware

var (
	stack middlewareStack
)

func Use(ms ...middleware) {
	stack = append(stack, ms...)
}

func Handle(path string, h Handler) {
	http.Handle(path, stack.run(h))
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res := NewResponse(w, r)
	h(res)
	res.Flush()
}

func (ms middlewareStack) run(h Handler) Handler {
	for i := len(ms); i > 0; i-- {
		h = ms[i-1](h)
	}
	return h
}

func Verb(method string, h Handler) Handler {
	return func(r *Response) {
		if r.Method == method {
			h(r)
		} else {
			MethodNotAllowed(r)
		}
	}
}

func Get(h Handler) Handler { return Verb("GET", h) }

func NotFound(r *Response) {
	r.Status = 404
	r.Body = "404 page not found"
}

func MethodNotAllowed(r *Response) {
	r.Status = 405
}
