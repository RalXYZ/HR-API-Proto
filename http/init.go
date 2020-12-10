package http

import (
	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func InitWebFramework() {
	e = echo.New()
	e.HideBanner = true
	addRoutes(e)
}

func StartServer() {
	e.Logger.Fatal(e.Start(":8080"))
}
