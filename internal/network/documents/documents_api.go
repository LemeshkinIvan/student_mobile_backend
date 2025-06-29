package doc_api

import (
	"github.com/gin-gonic/gin"
)

type DocumentsApi struct {
	Engine     *gin.Engine
	Controller *DocumentsController
	// log Log
}

func NewDocumentsApi(e *gin.Engine, c *DocumentsController) *DocumentsApi {
	api := &DocumentsApi{Engine: e, Controller: c}
	api.RegisterDocumentsRoutes()
	return api
}
