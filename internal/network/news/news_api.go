package news_api

import (
	"app/internal/bootstrap/db"
	l "app/internal/common/logger"
	"app/internal/data/repositories"

	"github.com/gin-gonic/gin"
)

type NewsApi struct {
	engine     *gin.Engine
	controller *NewsController
	repo       *repositories.NewsRepositoryImpl
}

func NewNewsApi(e *gin.Engine, db *db.Postgres) *NewsApi {
	r := repositories.NewNewsRepository(db)
	l.Logg.Info("news repository was init")

	ctrl := NewNewsController(r)
	l.Logg.Info("news contoller was init")

	api := &NewsApi{
		engine:     e,
		controller: ctrl,
		repo:       r,
	}

	api.RegRoutes()
	return api
}
