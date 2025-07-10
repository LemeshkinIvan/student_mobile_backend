package schedule_api

import (
	"app/internal/bootstrap/db"
	l "app/internal/common/logger"
	"app/internal/data/repositories"

	"github.com/gin-gonic/gin"
)

type ScheduleApi struct {
	engine     *gin.Engine
	controller *ScheduleController
	repo       *repositories.ScheduleRepositoryImpl
}

func NewScheduleApi(e *gin.Engine, db *db.Postgres) *ScheduleApi {
	r := repositories.NewScheduleRepositoryImpl(db)
	l.Logg.Info("schedule repository was init")

	ctrl := NewScheduleController(r)
	l.Logg.Info("schedule contoller was init")

	api := &ScheduleApi{
		engine:     e,
		controller: ctrl,
		repo:       r,
	}

	api.RegRoutes()
	return api
}
