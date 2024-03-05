package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/yassinebenaid/gomyadmin/database"
	"github.com/yassinebenaid/gomyadmin/database/repositories"
	"github.com/yassinebenaid/gomyadmin/http/resources"
)

func ListSchemas(ctx echo.Context) error {
	repository := repositories.SchemaRepotitory{
		Connection: database.Connection{
			Username: "root",
			Password: "password",
		},
	}

	schemas, err := repository.List()
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, resources.SchemaResource{
		Data:  schemas,
		Total: len(schemas),
	})
}
