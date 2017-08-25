package utils

import (
	"net/http"

	"github.com/golang/glog"
)

func Listen(l string, h http.Handler) {
	err := Setup()
	if err != nil {
		return
	}
	glog.Infoln("Listen on", l)
	glog.Fatal(http.ListenAndServe(l, h))
}
