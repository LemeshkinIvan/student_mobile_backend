package auth_api

import (
	"github.com/gin-gonic/gin"
)

type AuthApi struct {
	Engine     *gin.Engine
	Controller *AuthController
	// log Log
}

func NewAuthApi(e *gin.Engine, c *AuthController) *AuthApi {
	api := &AuthApi{Engine: e, Controller: c}
	api.RegisterAuthRoutes()
	return api
}
