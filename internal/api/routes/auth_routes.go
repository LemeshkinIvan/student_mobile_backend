package routes

func (r *RouteEntity) RegisterAuthRoutes() {
	authRoutes := r.Engine.Group("/auth")

	{
		authRoutes.POST("/token", r.AuthController.GetTokens)
		authRoutes.POST("/getUser", r.AuthController.GetUsers)
		authRoutes.POST("/refreshTokens", r.AuthController.RefreshToken)
	}
}
