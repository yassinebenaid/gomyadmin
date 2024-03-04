package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yassinebenaid/gomyadmin/http/handlers"
)

type Route struct {
	Method  string
	Path    string
	Handler echo.HandlerFunc
}

var DefaultMiddlewares = []echo.MiddlewareFunc{}

var Routes = []Route{
	{"GET", "databases", handlers.ListDatabases},
}
