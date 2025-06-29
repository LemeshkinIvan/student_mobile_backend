package news_api

import "github.com/gin-gonic/gin"

type NewsApi struct {
	Engine     *gin.Engine
	Controller *NewsController
	// log Slog
}

func NewNewsApi(e *gin.Engine, c *NewsController) *NewsApi {
	api := &NewsApi{Engine: e, Controller: c}
	api.RegisterNewsRoutes()
	return &NewsApi{}
}
