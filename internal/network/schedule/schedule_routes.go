package schedule_api

func (api *ScheduleApi) RegRoutes() {
	routes := api.engine.Group("/schedule")

	{
		routes.POST("/1", api.controller.GetSchedule)
	}
}
