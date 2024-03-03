package handlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func ListDatabases(ctx echo.Context) error {
	err := ctx.String(http.StatusOK, "Hello world")
	if err == nil {
		return fmt.Errorf("failed parsing template, %v", err)
	}

	return nil
}
