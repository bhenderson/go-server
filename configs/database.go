package configs

import (
	"os"

	"github.com/bhenderson/go-server/utils"
)

var DB = DBConfig{}

type DBConfig struct {
	Adapter string
	Timeout int
	Pool    int

	Url string

	DBName                  string
	User                    string
	password                string
	Host                    string
	Port                    string
	SSLMode                 string
	FallbackApplicationName string
	ConnectTimeout          int
	SSLCert                 string
	SSLKey                  string
	SSLRootCert             string
}

var dbConnectTemplate = `
{{- if .DBName}}dbname={{.DBName}} {{end -}}
{{- if .User}}user={{.User}} {{end -}}
{{- if .Password}}password='{{.Password | printf "%s"}}' {{end -}}
{{- if .Host}}host={{.Host}} {{end -}}
{{- if .Port}}port={{.Port}} {{end -}}
{{- if .SSLMode}}sslmode={{.SSLMode}} {{end -}}
{{- if .FallbackApplicationName}}fallback_application_name={{.FallbackApplicationName}} {{end -}}
{{- if .ConnectTimeout}}connect_timeout={{.ConnectTimeout}} {{end -}}
{{- if .SSLCert}}sslcert={{.SSLCert}} {{end -}}
{{- if .SSLKey}}sslkey={{.SSLKey}} {{end -}}
{{- if .SSLRootCert}}sslrootcert={{.SSLRootCert}} {{end -}}
`

func (c DBConfig) Connect() string {
	if c.Url != "" {
		return c.Url
	}
	return utils.MustTemplateString(dbConnectTemplate, c)
}

func (c DBConfig) Password() string {
	return c.password
}

func init() {
	All["db"] = &DB

	utils.Setup(func() error {
		if u, ok := os.LookupEnv("DATABASE_URL"); ok {
			DB.Url = u
			return nil
		}

		DB.Adapter = "postgres"
		DB.Host = "localhost"
		DB.User = "user"
		DB.password = "user"
		DB.Port = os.Getenv("PGPORT")

		switch {
		case Env.IsProd():
		case Env.IsStage():
		case Env.IsTest():
		case Env.IsDev():
		}

		return nil
	})
}
