package auth_api

func (api *AuthApi) RegRoutes() {
	authRoutes := api.engine.Group("/auth")

	{
		authRoutes.POST("/token", api.controller.GetTokens)
		authRoutes.POST("/getUser", api.controller.GetUsers)
		authRoutes.POST("/refreshTokens", api.controller.RefreshToken)
	}
}
