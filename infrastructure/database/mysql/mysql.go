// Package mysql provides ...
package mysql

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"iohttps.com/live/realworld-medium-rewrite/infrastructure/database"
)

//MysqlHandler handler of mysql
type mysqlHandler struct {
	Conn *sql.DB
}

//NewMysqlHandler new mysqlHandler
func NewMysqlHandler(dataSourceName string) (database.DbHandler, error) {
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(time.Duration(4 * time.Hour))
	return &mysqlHandler{db}, err
}

//Execute .....
func (handler *mysqlHandler) Execute(query string, args ...interface{}) (database.Result, error) {
	result, err := handler.Conn.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	return &mysqlResult{result}, nil
}

//Query .....
func (handler *mysqlHandler) Query(query string, args ...interface{}) (database.Rows, error) {
	rows, err := handler.Conn.Query(query, args...)
	if err != nil {
		return nil, err
	}

	return &mysqlRows{rows}, nil
}

//QueryRow ....
func (handler *mysqlHandler) QueryRow(query string, args ...interface{}) database.Row {
	row := handler.Conn.QueryRow(query, args...)
	return &mysqlRow{row}
}

//Ping ....
func (handler *mysqlHandler) Ping() error {
	return handler.Conn.Ping()
}

//PingContext ....
func (handler *mysqlHandler) PingContext(ctx context.Context) error {
	return handler.Conn.PingContext(ctx)
}

//mysqlResult result
type mysqlResult struct {
	Result sql.Result
}

//LastInsertId ....
func (result *mysqlResult) LastInsertId() (int64, error) {
	return result.Result.LastInsertId()
}

//RowsAffected ....
func (result *mysqlResult) RowsAffected() (int64, error) {
	return result.Result.RowsAffected()
}

//mysqlRows rows
type mysqlRows struct {
	Rows *sql.Rows
}

//Scan ....
func (rows *mysqlRows) Scan(dest ...interface{}) error {
	return rows.Rows.Scan(dest...)
}

//Next ...
func (rows *mysqlRows) Next() bool {
	return rows.Rows.Next()
}

type mysqlRow struct {
	Row *sql.Row
}

//Scan ...
func (row *mysqlRow) Scan(dest ...interface{}) error {
	return row.Row.Scan(dest...)
}
