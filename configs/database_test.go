package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBConnect(t *testing.T) {
	db := DBConfig{}
	assert.Equal(t, "", db.Connect())

	db.Host = "localhost"
	db.password = "foo bar"
	assert.Equal(t, `password="foo bar" host=localhost `, db.Connect())
}
