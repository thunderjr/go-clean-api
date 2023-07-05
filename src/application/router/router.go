package router

import "github.com/labstack/echo/v4"

type Router struct {
	server *echo.Echo
}

func Init(server *echo.Echo) error {
	var router = &Router{server: server}

	CreateUserRoute(router)

	return nil
}
