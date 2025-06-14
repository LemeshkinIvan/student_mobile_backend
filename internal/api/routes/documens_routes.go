package routes

func (r *RouteEntity) RegisterDocumentsRoutes() {
	documentsRoutes := r.Engine.Group("/documents")

	{
		documentsRoutes.POST("/byUid", r.AuthController.GetTokens)
		documentsRoutes.POST("/all", r.AuthController.GetUsers)
		documentsRoutes.POST("/user/byUid", r.AuthController.RefreshToken)
	}
}
