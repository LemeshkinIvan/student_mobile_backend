package routes

import (
	"app/internal/api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouteEntity struct {
	Engine         *gin.Engine
	AuthController *controllers.AuthController
}

func NewRouteEntity(authController *controllers.AuthController) *RouteEntity {
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
}
