package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yassinebenaid/gomyadmin/http/handlers"
)

func Register(echo *echo.Echo) {
	api := echo.Group("/api/*")

	api.GET("databases", handlers.ListDatabases)
}
