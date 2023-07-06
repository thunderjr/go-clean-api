package http

import (
	"reflect"

	"github.com/labstack/echo/v4"
	"github.com/thunderjr/go-clean-api/src/application/helpers"
	"github.com/thunderjr/go-clean-api/src/application/router"
	"github.com/thunderjr/go-clean-api/src/presentation/controllers"
)

type EchoRouter struct {
	e *echo.Echo
}

func NewEchoRouter(e *echo.Echo) router.Router {
	return &EchoRouter{
		e: e,
	}
}

func (r *EchoRouter) Handle(method string, path string, controller controllers.Controller) error {
	return helpers.InvokeWithError(r.e, method, path, func(ctx echo.Context) error {
		var data *interface{} = new(interface{})
		if err := ctx.Bind(&data); err != nil {
			return err
		}

		res := controller.Handle(reflect.ValueOf(*data).Interface().(map[string]interface{}))
		return ctx.JSON(res.Status, res.Data)
	})
}

func (r *EchoRouter) GET(path string, controller controllers.Controller) error {
	return r.Handle("GET", path, controller)
}

func (r *EchoRouter) POST(path string, controller controllers.Controller) error {
	return r.Handle("POST", path, controller)
}

func (r *EchoRouter) PUT(path string, controller controllers.Controller) error {
	return r.Handle("PUT", path, controller)
}

func (r *EchoRouter) DELETE(path string, controller controllers.Controller) error {
	return r.Handle("DELETE", path, controller)
}
