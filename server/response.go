package server

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"
)

type Response struct {
	Status int
	header http.Header
	Body   interface{}

	Time time.Time

	w http.ResponseWriter
	*http.Request
	flushed       bool
	contentLength int
}

func NewResponse(w http.ResponseWriter, r *http.Request) *Response {
	return &Response{
		Time:    time.Now(),
		w:       w,
		Request: r,
	}
}

func (r *Response) WriteHeader(s int) {
	r.Status = s
}

func (r *Response) Header() http.Header {
	return r.w.Header()
}

func (r *Response) Write(p []byte) (int, error) {
	if x, ok := r.Body.([]byte); ok {
		x = append(x, p...)
		return len(x), nil
	}
	if r.Body == nil {
		r.Body = p
		return len(p), nil
	}
	return 0, errors.WithStack(bodyAssignedError{})
}

func (r *Response) Flush() {
	if r.flushed {
		return
	}
	r.flushed = true

	w := r.w
	w.WriteHeader(r.Status)
	b := r.Body
	if b == nil {
		b = r.Status
	}
	var i64 int64
	var i int

	switch x := b.(type) {
	case io.WriterTo:
		i64, _ = x.WriteTo(w)
	case int:
		i, _ = io.WriteString(w, http.StatusText(x))
	case string:
		i, _ = io.WriteString(w, x)
	case []byte:
		i, _ = w.Write(x)
	default:
		i, _ = fmt.Fprintf(w, "%v", x)
	}

	if i64 > 0 {
		i = int(i64)
	}
	r.contentLength += i

	return
}

const dash = "-"

func (r *Response) RemoteAddr() string {
	host, _, err := net.SplitHostPort(r.Request.RemoteAddr)
	if err != nil {
		return dash
	}
	return host
}

// Username returns the Username or a "-"
func (r *Response) Username() string {
	if r.URL != nil {
		if u := r.URL.User; u != nil {
			return u.Username()
		}
	}
	return dash
}

func (r *Response) LocalTime() string {
	return r.Time.Format("02/Jan/2006:15:04:05 -0700")
}

func (r *Response) RequestLine() string {
	return fmt.Sprintf("%s %s %s", r.Method, r.RequestURI, r.Proto)
}

func (r *Response) ContentLength() int {
	return r.contentLength
}

func (r *Response) ContentSize() string {
	if r.contentLength == 0 {
		return dash
	}
	return strconv.Itoa(r.contentLength)
}

func (r *Response) Since() time.Duration {
	return time.Since(r.Time)
}
