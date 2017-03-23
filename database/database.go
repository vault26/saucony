package database

import (
	"github.com/golang/glog"
	pg "gopkg.in/pg.v5"
)

type DB struct {
	*pg.DB
}

func logError(err error) {
	if err != nil {
		glog.Error(err)
	}
}
