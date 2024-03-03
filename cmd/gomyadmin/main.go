package main

import (
	"github.com/labstack/echo/v4"
	"github.com/yassinebenaid/gomyadmin/http/routes"
	"github.com/yassinebenaid/gomyadmin/web"
)

func main() {

	e := echo.New()

	routes.Register(e)

	e.FileFS("/*", "dist/index.html", web.StaticFS)
	e.FileFS("/favicon.ico", "dist/favicon.ico", web.StaticFS)
	e.StaticFS("/assets/*", echo.MustSubFS(web.StaticFS, "dist/assets"))

	e.Logger.Fatal(e.Start(":7000"))
}
