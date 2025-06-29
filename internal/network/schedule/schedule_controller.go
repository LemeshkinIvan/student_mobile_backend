package schedule_api

import (
	r "app/internal/data/repositories"

	"github.com/gin-gonic/gin"
)

type ScheduleController struct {
	Repository *r.ScheduleRepositoryImpl
}

func NewScheduleController(r *r.ScheduleRepositoryImpl) *ScheduleController {
	return &ScheduleController{
		Repository: r,
	}
}

func (ctrl *ScheduleController) GetSchedule(c *gin.Context) {

}
