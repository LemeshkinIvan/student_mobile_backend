package news_api

func (api *NewsApi) RegisterNewsRoutes() {
	routes := api.Engine.Group("/schedule")

	{
		routes.GET("/2", api.Controller.GetNews)
	}
}
