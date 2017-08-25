package endpoints

import (
	"github.com/bhenderson/go-server/configs"
	"github.com/bhenderson/go-server/server"
	"github.com/bhenderson/go-server/utils"
)

func init() {
	utils.Setup(func() error {
		server.Use(server.LogHandler)
		server.Handle("/", server.NotFound)
		server.Handle("/foo", index)

		server.Use(server.JSONHandler)

		server.Handle("/_priv/config", configHandler)
		server.Handle("/api/foo", server.Get(foo))
		return nil
	})
}

func configHandler(r *server.Response) {
	r.Body = configs.All
}

func index(r *server.Response) {
	r.Body = "hello world"
}

func foo(r *server.Response) {
	r.Body = map[string]string{"hello": "world"}
}
