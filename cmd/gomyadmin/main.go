package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yassinebenaid/gomyadmin/http/routes"
	"github.com/yassinebenaid/gomyadmin/web"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())

	api := e.Group("/api/")
	api.Use(routes.Middlewares...)

	for _, r := range routes.Routes {
		api.Add(r.Method, r.Path, r.Handler, r.Middlewares...)
	}

	e.FileFS("/*", "dist/index.html", web.StaticFS)
	e.FileFS("/favicon.ico", "dist/favicon.ico", web.StaticFS)
	e.StaticFS("/assets/*", echo.MustSubFS(web.StaticFS, "dist/assets"))

	e.Logger.Fatal(e.Start(":7000"))
}
