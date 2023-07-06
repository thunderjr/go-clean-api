package router

import (
	"github.com/thunderjr/go-clean-api/src/presentation/controllers"
)

type Router interface {
	Handle(method string, path string, handler controllers.Controller) error
	GET(path string, handler controllers.Controller) error
	POST(path string, handler controllers.Controller) error
	PUT(path string, handler controllers.Controller) error
	DELETE(path string, handler controllers.Controller) error
}
