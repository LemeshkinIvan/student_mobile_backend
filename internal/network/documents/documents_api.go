package doc_api

import (
	"app/internal/bootstrap/db"
	l "app/internal/common/logger"
	"app/internal/data/repositories"

	"github.com/gin-gonic/gin"
)

type DocumentsApi struct {
	engine     *gin.Engine
	controller *DocumentsController
	repo       *repositories.DocumentsRepositoryImpl
}

func NewDocumentsApi(e *gin.Engine, db *db.Postgres) *DocumentsApi {
	r := repositories.NewDocumentsRepository(db)
	l.Logg.Info("doc repository was init")

	ctrl := NewDocumentsController(r)
	l.Logg.Info("doc contoller was init")

	api := &DocumentsApi{
		engine:     e,
		controller: ctrl,
		repo:       r,
	}

	api.RegRoutes()
	return api
}
