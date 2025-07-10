package doc_api

func (api *DocumentsApi) RegRoutes() {
	documentsRoutes := api.engine.Group("/documents")

	{
		documentsRoutes.POST("/byUid", api.controller.GetAllDocuments)
		documentsRoutes.POST("/all", api.controller.GetDocumentByUid)
		documentsRoutes.POST("/user/byUid", api.controller.GetUserDocuments)
	}
}
