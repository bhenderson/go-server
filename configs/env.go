package configs

import "github.com/bhenderson/go-server/utils"

var Env env

func init() {
	SetEnv("dev")
	All["env"] = &Env
}

type env string

func (e env) IsProd() bool  { return e == "prod" }
func (e env) IsStage() bool { return e == "stage" }
func (e env) IsTest() bool  { return e == "test" }
func (e env) IsDev() bool   { return e == "dev" }

func SetEnv(e string) {
	Env = env(utils.LookupEnvDefault("GO_ENV", e))
}
