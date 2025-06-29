package auth_api

func (api *AuthApi) RegisterAuthRoutes() {
	authRoutes := api.Engine.Group("/auth")

	{
		authRoutes.POST("/token", api.Controller.GetTokens)
		authRoutes.POST("/getUser", api.Controller.GetUsers)
		authRoutes.POST("/refreshTokens", api.Controller.RefreshToken)
	}
}
