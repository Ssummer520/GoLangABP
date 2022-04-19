package datasource

import "github.com/jmoiron/sqlx"

type IDb interface {
	DB() *sqlx.DB
	Connect() error
}
