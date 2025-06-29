package schedule_api

import (
	"github.com/gin-gonic/gin"
)

type ScheduleApi struct {
	Engine     *gin.Engine
	Controller *ScheduleController
	// log Log
}

func NewScheduleApi(e *gin.Engine, c *ScheduleController) *ScheduleApi {
	api := &ScheduleApi{Engine: e, Controller: c}
	api.RegisterScheduleRoutes()
	return api
}
