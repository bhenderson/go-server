package db

import (
	"database/sql"

	"github.com/golang/glog"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
	"github.com/bhenderson/go-server/configs"
	"github.com/bhenderson/go-server/utils"
)

var db *sql.DB

func Connect() error {
	c := configs.DB

	var err error

	adapter, connString := c.Adapter, c.Connect()
	glog.V(2).Infoln("connecting to db", adapter, connString)
	db, err = sql.Open(adapter, connString)
	if err != nil {
		return errors.Wrap(err, "sql open")
	}

	err = db.Ping()
	if err != nil {
		return errors.Wrap(err, "sql ping")
	}
	return nil
}

func Close() error {
	if db == nil {
		return nil
	}
	return errors.Wrap(db.Close(), "database close")
}

func init() {
	utils.Setup(Connect)
	utils.Teardown(Close)
}
