package configs

import "github.com/bhenderson/go-server/utils"

var Listen = utils.LookupEnvDefault("LISTEN", ":8080")

var All = make(map[string]interface{})
