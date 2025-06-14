package bootstrap

import (
	auth_ctrl "app/internal/api/controllers/auth"
	doc_ctrl "app/internal/api/controllers/documents"
	"app/internal/api/routes"
	documents_repository "app/internal/data/repository/documents"
	user_repository "app/internal/data/repository/user"
	doc_serv "app/internal/services/document_service"
	user_serv "app/internal/services/user_service"
)

type App struct {
	Router *routes.RouteEntity
	Db     *DatabaseManager
	Env    *AppEnvironment
}

func InitApp() *App {
	userService := user_serv.NewUserService()
	documentsService := doc_serv.NewDocumentsService()

	authController := auth_ctrl.NewAuthController(
		userService, user_repository.NewUserRepository(),
	)

	documentsController := doc_ctrl.NewDocumentsController(
		documentsService, documents_repository.NewDocumentsRepository(),
	)

	router := routes.NewRouteEntity(
		authController,
		documentsController,
	)

	router.RegisterAllRoutes()

	database := NewDatabaseManager()
	env := NewAppEnvironment()

	return &App{
		Router: router,
		Db:     database,
		Env:    env,
	}
}

// for local debug
func (a *App) Run() error {
	a.Router.RegisterAllRoutes()
	return a.Router.Engine.Run("fck_all")
}
