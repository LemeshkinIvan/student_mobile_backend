package news_api

func (api *NewsApi) RegRoutes() {
	routes := api.engine.Group("/schedule")

	{
		routes.GET("/2", api.controller.GetNews)
	}
}
