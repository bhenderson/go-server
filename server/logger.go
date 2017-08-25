package server

import (
	"github.com/golang/glog"
	"github.com/bhenderson/go-server/utils"
)

var (
	common   = `{{.RemoteAddr}} - {{.Username}} [{{.LocalTime}}] "{{.RequestLine}}" {{.Status}} {{.ContentSize}}`
	combined = common + ` "{{.Referer}}" "{{.UserAgent}}"`
	custom   = combined + ` ({{.Since}})`
)

func LogHandler(next Handler) Handler {
	tmpl := utils.MustTemplate(custom)
	return func(r *Response) {
		next(r)

		r.Flush()

		s, err := utils.TemplateExecute(tmpl, r)
		if err != nil {
			glog.Error(r, err)
		} else {
			glog.Infoln(s)
		}
	}
}
