package utils

import (
	"bytes"
	"os"
	"text/template"

	"github.com/pkg/errors"
)

func MustTemplate(tmpl string) *template.Template {
	return template.Must(Template(tmpl))
}

func MustTemplateString(tmpl string, v interface{}) string {
	s, err := TemplateString(tmpl, v)
	if err != nil {
		panic(err)
	}
	return s
}

func TemplateString(tmpl string, v interface{}) (string, error) {
	t, err := Template(tmpl)
	if err != nil {
		return "", err
	}
	s, err := TemplateExecute(t, v)
	if err != nil {
		return "", err
	}
	return s, nil
}

func Template(tmpl string) (*template.Template, error) {
	t, err := template.New("utils").Parse(tmpl)
	if err != nil {
		err = errors.Wrap(err, "template parse failed")
	}
	return t, err
}

func TemplateExecute(t *template.Template, v interface{}) (string, error) {
	var buf bytes.Buffer
	err := t.Execute(&buf, v)
	if err != nil {
		return "", errors.Wrap(err, "template execute failed")
	}
	return buf.String(), nil
}

func LookupEnvDefault(name, def string) string {
	if s, ok := os.LookupEnv(name); ok {
		return s
	}
	return def
}
