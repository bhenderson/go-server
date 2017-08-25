package main

import (
	"flag"

	"github.com/bhenderson/go-server/configs"
	_ "github.com/bhenderson/go-server/db"
	_ "github.com/bhenderson/go-server/endpoints"
	"github.com/bhenderson/go-server/utils"
)

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()

	defer utils.Teardown()

	utils.Listen(configs.Listen, nil)
}
