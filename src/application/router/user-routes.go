package router

import (
	"github.com/labstack/echo/v4"
	"github.com/thunderjr/go-clean-api/src/application/features"
	"github.com/thunderjr/go-clean-api/src/domain/entities"
	local_database "github.com/thunderjr/go-clean-api/src/infra/database/local"
	local_repositories "github.com/thunderjr/go-clean-api/src/infra/database/local/repositories"
	"github.com/thunderjr/go-clean-api/src/presentation/controllers"
)

func CreateUserRoute(router *Router) {
	repository := local_repositories.GetLocalUserRepository(local_database.DB)
	feature := features.GetCreateUserFeature(repository)
	controller := controllers.GetCreateUserController(feature)

	router.server.POST("/user", func(ctx echo.Context) error {
		data := new(entities.User)
		if err := ctx.Bind(data); err != nil {
			return err
		}

		res := controller.Handle(data)
		return ctx.JSON(res.Status, res.Data)
	})
}
