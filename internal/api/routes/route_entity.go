package routes

import (
	user_ctrl "app/internal/api/controllers/auth"
	documents_ctrl "app/internal/api/controllers/documents"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouteEntity struct {
	Engine              *gin.Engine
	AuthController      *user_ctrl.AuthController
	DocumentsController *documents_ctrl.DocumentsController
}

func NewRouteEntity(
	authController *user_ctrl.AuthController,
	documentsController *documents_ctrl.DocumentsController,
) *RouteEntity {
	engine := gin.Default()

	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	return &RouteEntity{
		Engine:         engine,
		AuthController: authController,
	}
}

func (re *RouteEntity) RegisterAllRoutes() {
	re.RegisterAuthRoutes()
	re.RegisterDocumentsRoutes()
}
