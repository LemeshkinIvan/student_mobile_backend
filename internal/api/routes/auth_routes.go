package routes

func (r *RouteEntity) RegisterAuthRoutes() {
	albumRoutes := r.Engine.Group("/auth")

	{
		albumRoutes.GET("/token", r.AuthController.GetTokens)
		albumRoutes.GET("/getUser", r.AuthController.GetUsers)
		albumRoutes.GET("/refreshTokens", r.AuthController.RefreshToken)
	}
}
