package doc_api

func (api *DocumentsApi) RegisterDocumentsRoutes() {
	documentsRoutes := api.Engine.Group("/documents")

	{
		documentsRoutes.POST("/byUid", api.Controller.GetAllDocuments)
		documentsRoutes.POST("/all", api.Controller.GetDocumentByUid)
		documentsRoutes.POST("/user/byUid", api.Controller.GetUserDocuments)
	}
}
