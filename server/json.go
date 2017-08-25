package server

import (
	"encoding/json"

	"github.com/pkg/errors"
)

func JSONHandler(next Handler) Handler {
	return func(r *Response) {
		next(r)

		r.Header().Set("Content-Type", "application/json")
		b, err := json.Marshal(r.Body)
		if err != nil {
			r.Status = 500
			r.Body = errors.Wrap(err, "json marshal failed")
		} else {
			r.Body = b
		}
	}
}
