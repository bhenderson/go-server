package utils

import "github.com/golang/glog"

type runFunc func() error

var (
	setupFuncs    []runFunc
	teardownFuncs []runFunc
)

func Setup(fs ...runFunc) error    { return run(&setupFuncs, fs...) }
func Teardown(fs ...runFunc) error { return run(&teardownFuncs, fs...) }

func run(funcs *[]runFunc, fs ...runFunc) error {
	if fs != nil {
		*funcs = append(*funcs, fs...)
		return nil
	}

	for _, f := range *funcs {
		if e := f(); e != nil {
			glog.Error(e)
			return e
		}
	}

	return nil
}
