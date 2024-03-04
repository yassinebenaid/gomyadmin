package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yassinebenaid/gomyadmin/database"
)

func ListDatabases(ctx echo.Context) error {
	conn, err := database.Connect(database.Connection{
		Username: "root",
		Password: "password",
	})
	if err != nil {
		return fmt.Errorf("database connection error : %v", err)
	}

	r, err := conn.Query("show databases")
	if err != nil {
		return fmt.Errorf("database query error : %v", err)
	}

	var databases []string

	for r.Next() {
		var db_name string
		if err := r.Scan(&db_name); err != nil {
			return fmt.Errorf("database read error : %v", err)
		}
		databases = append(databases, db_name)
	}

	return ctx.JSON(http.StatusOK, databases)
}
