package database

import pg "gopkg.in/pg.v5"

type DB struct {
	*pg.DB
}
