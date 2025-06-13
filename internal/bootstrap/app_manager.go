package bootstrap

import (
	"app/internal/api/controllers"
	"app/internal/api/routes"
	user_repository "app/internal/data/repository"
	"app/internal/services"
)

type App struct {
	Router *routes.RouteEntity
	Db     *DatabaseManager
	Env    *AppEnvironment
}

func InitApp() *App {
	userService := services.NewUserService()

	authController := controllers.NewAuthController(
		userService, user_repository.NewUserRepository(),
	)

	router := routes.NewRouteEntity(authController)

	database := NewDatabaseManager()
	env := NewAppEnvironment()

	return &App{
		Router: router,
		Db:     database,
		Env:    env,
	}
}

func (a *App) Run() error {
	a.Router.RegisterAllRoutes()
	return a.Router.Engine.Run(":8080")
}
