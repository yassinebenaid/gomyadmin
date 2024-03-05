package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yassinebenaid/gomyadmin/database/repositories"
)

func ListSchemas(ctx echo.Context) error {
	conn, ok := ctx.Get("db:connection").(*sql.DB)
	if !ok {
		return fmt.Errorf("context missing database connection")
	}
	repository := repositories.SchemaRepotitory{Connection: conn}

	schemas, err := repository.ListSchemas()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusOK, schemas)
}

func ListCollations(ctx echo.Context) error {
	conn, ok := ctx.Get("db:connection").(*sql.DB)
	if !ok {
		return fmt.Errorf("context missing database connection")
	}
	repository := repositories.SchemaRepotitory{Connection: conn}

	collations, err := repository.ListCollations()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, collations)
}
