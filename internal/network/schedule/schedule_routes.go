package schedule_api

func (api *ScheduleApi) RegisterScheduleRoutes() {
	routes := api.Engine.Group("/schedule")

	{
		routes.POST("/1", api.Controller.GetSchedule)
	}
}
