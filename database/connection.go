package database

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const driver = "mysql"

type Connection struct {
	Username string
	Password string
	Protocol string
	Address  string
	DBName   string
	Params   map[string]string
}

var DefaultConnection = Connection{}

func (c Connection) DSN() string {
	var dsn string

	dsn += c.Username

	if c.Username != "" && c.Password != "" {
		dsn += ":" + c.Password
	}

	dsn += "@"
	if c.Address != "" {
		if c.Protocol == "" {
			dsn += "tcp"
		} else {
			dsn += c.Protocol
		}
		dsn += "(" + c.Address + ")"
	}

	dsn += "/" + c.DBName
	if len(c.Params) > 0 {
		dsn += "?"
	}
	var delimiter string
	for key, value := range c.Params {
		dsn += delimiter + key + "=" + value
		delimiter = "&"
	}

	return dsn
}

func Connect(c Connection) (*sql.DB, error) {
	conn, err := sql.Open(driver, c.DSN())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to [%s] server, %v", driver, err)
	}

	if err := conn.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to [%s] server, %v", driver, err)
	}

	return conn, nil
}
