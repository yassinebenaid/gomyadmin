package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yassinebenaid/gomyadmin/database"
	"github.com/yassinebenaid/gomyadmin/database/repositories"
	"github.com/yassinebenaid/gomyadmin/http/resources"
)

func ListDatabases(ctx echo.Context) error {
	repository := repositories.DatabaseRepotitory{
		Connection: database.Connection{
			Username: "root",
			Password: "password",
		},
	}

	databases, err := repository.List()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resources.DatabaseResource{
		Data:  databases,
		Total: len(databases),
	})
}
