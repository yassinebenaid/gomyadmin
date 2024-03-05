package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yassinebenaid/gomyadmin/database"
	"github.com/yassinebenaid/gomyadmin/database/repositories"
)

func ListSchemas(ctx echo.Context) error {
	repository := repositories.SchemaRepotitory{
		Connection: database.Connection{
			Username: "root",
			Password: "password",
		},
	}

	schemas, err := repository.SchemasList()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, schemas)
}

func ListCollations(ctx echo.Context) error {
	repository := repositories.SchemaRepotitory{
		Connection: database.Connection{
			Username: "root",
			Password: "password",
		},
	}

	collations, err := repository.CollationsList()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, collations)
}
