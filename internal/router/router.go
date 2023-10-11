package router

import (
	"echo-mangosteen/internal/controller"
	"echo-mangosteen/pkg/jwt"
	"echo-mangosteen/pkg/middleware"

	"github.com/labstack/echo/v4"
)

func NewRouter(
	jwt *jwt.JWT,
	pingCtrl controller.PingController,
	userCtrl controller.UserController,
	codeCtrl controller.CodeController,
	tagCtrl controller.TagController,
) *echo.Echo {
	e := echo.New()

	authGroup := e.Group("/")
	authGroup.Use(middleware.JWTMiddleware(jwt))
	{
		authGroup.POST("tags/add", tagCtrl.AddTag)
		authGroup.GET("/ping", pingCtrl.Ping)
	}

	guestGroup := e.Group("")
	{
		guestGroup.POST("/login", userCtrl.Login)
	}

	return e
}
