package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/yassinebenaid/gomyadmin/http/handlers"
	"github.com/yassinebenaid/gomyadmin/http/middleware"
)

var Middlewares = []echo.MiddlewareFunc{
	middleware.ConnectToDatabase,
}

type Route struct {
	Method      string
	Path        string
	Handler     echo.HandlerFunc
	Middlewares []echo.MiddlewareFunc
}

var Routes = []Route{
	{"GET", "schemas", handlers.ListSchemas, nil},
	{"POST", "schemas", handlers.CreateSchema, nil},
	{"GET", "collations", handlers.ListCollations, nil},
}
