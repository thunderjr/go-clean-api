package router

func InitRoutes(router *Router) {
	CreateUserRoute(router)
	GetUserRoute(router)
}
