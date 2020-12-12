// Package database provides ...
package database

import "context"

//DbHandler Handler of db
type DbHandler interface {
	Execute(query string, args ...interface{}) (Result, error)
	Query(query string, args ...interface{}) (Rows, error)
	QueryRow(query string, args ...interface{}) Row
	Ping() error
	PingContext(ctx context.Context) error
	Begin() (Tx, error)
}

//Result ....
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

//Rows ....
type Rows interface {
	Scan(dest ...interface{}) error
	Next() bool
}

//Row ....
type Row interface {
	Scan(dest ...interface{}) error
}

type Tx interface {
}

//DbRepo ....
type DbRepo struct {
	Handler DbHandler
}
