package router

import (
	"github.com/thunderjr/go-clean-api/src/application/features"
	local_database "github.com/thunderjr/go-clean-api/src/infra/database/local"
	local_repositories "github.com/thunderjr/go-clean-api/src/infra/database/local/repositories"
	"github.com/thunderjr/go-clean-api/src/presentation/controllers"
)

func CreateUserRoute(router *Router) {
	repository := local_repositories.NewLocalUserRepository(local_database.DB)
	feature := features.NewCreateUserFeature(repository)
	controller := controllers.NewCreateUserController(feature)

	(*router).Handle("POST", "/user", controller)
}

func GetUserRoute(router *Router) {
	repository := local_repositories.NewLocalUserRepository(local_database.DB)
	feature := features.NewGetUserFeature(repository)
	controller := controllers.NewGetUserController(feature)

	(*router).Handle("GET", "/user", controller)
}
