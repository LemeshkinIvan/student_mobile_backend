package auth_api

import (
	db "app/internal/bootstrap/db"
	l "app/internal/common/logger"
	repo "app/internal/data/repositories"

	"github.com/gin-gonic/gin"
)

type AuthApi struct {
	engine     *gin.Engine
	controller *AuthController
	repo       *repo.AuthRepositoryImpl
}

func NewAuthApi(g *gin.Engine, db *db.Postgres) *AuthApi {
	r := repo.NewAuthRepository(db)
	l.Logg.Info("auth repository was init")

	ctrl := NewAuthController(r)
	l.Logg.Info("auth contoller was init")

	api := &AuthApi{
		engine:     g,
		controller: ctrl,
		repo:       r,
	}

	api.RegRoutes()
	return api
}
