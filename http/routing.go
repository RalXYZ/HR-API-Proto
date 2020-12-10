package http

import (
	"HR-API-proto/http/controllers"
	"github.com/labstack/echo/v4"
)

func addRoutes(e *echo.Echo) {
	api := e.Group("/api")
	api.POST("/login", controllers.Login)
	api.POST("/register", controllers.Register)
	api.GET("/member", controllers.GetMember)
	api.POST("/member", controllers.SetMember)
	api.DELETE("/member", controllers.DeleteMember)

	urlapi := e.Group("/urlapi")
	urlapi.GET("/login", controllers.LoginUrl)
	urlapi.GET("/register", controllers.RegisterUrl)
	urlapi.GET("/member", controllers.GetMemberUrl)
	urlapi.GET("/member/create", controllers.SetMemberUrl)
	urlapi.GET("/member/delete", controllers.DeleteMemberUrl)
}
