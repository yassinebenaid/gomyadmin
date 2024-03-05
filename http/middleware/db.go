package middleware

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/yassinebenaid/gomyadmin/database"
)

func ConnectToDatabase(next echo.HandlerFunc) echo.HandlerFunc {
	conn := database.Connection{
		Username: "root",
		Password: "password",
	}

	db_conn, err := database.Connect(conn)
	if err != nil {
		return func(c echo.Context) error {
			return fmt.Errorf("database connection error : %v", err)
		}
	}

	return func(c echo.Context) error {
		c.Set("db:connection", db_conn)
		return next(c)
	}
}
