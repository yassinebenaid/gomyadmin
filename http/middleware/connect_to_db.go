package middleware

import "github.com/labstack/echo/v4"

var s echo.MiddlewareFunc = func(next echo.HandlerFunc) echo.HandlerFunc {

	return next
}

// func ConnectToDB()
