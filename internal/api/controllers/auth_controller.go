package controllers

import (
	user_repository "app/internal/data/repository"
	auth_models "app/internal/domain/models/auth"
	"app/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	UserService *services.UserService
	Repository  *user_repository.UserRepositoryImpl
}

func NewAuthController(
	userService *services.UserService,
	repository *user_repository.UserRepositoryImpl,
) *AuthController {
	return &AuthController{
		UserService: userService,
		Repository:  repository,
	}
}

func (ctrl *AuthController) GetUsers(c *gin.Context) {
	var request auth_models.UserRequest

	if err := c.BindJSON(&request); err != nil {
		var alert string = "bad client getUser request " + err.Error()
		log.Panic(alert)

		c.JSON(http.StatusBadRequest, ErrorRessponse{
			Message: err.Error(),
		})
	}

	var user = ctrl.Repository.GetUserByAccessToken(request.Code)

	c.IndentedJSON(http.StatusOK, user)
}

func (ctrl *AuthController) GetTokens(c *gin.Context) {
	var tokenRequest auth_models.TokensPairRequest

	if err := c.BindJSON(&tokenRequest); err != nil {
		var alert string = "bad client token request " + err.Error()
		log.Panic(alert)

		c.JSON(http.StatusBadRequest, ErrorRessponse{
			Message: err.Error(),
		})

		return
	}

	var pair, err = ctrl.Repository.GetTokenPair(tokenRequest.Code)

	if err != nil {
		var alert string = "cant generate tokens " + err.Error()
		log.Panic(alert)

		c.JSON(http.StatusInternalServerError, "tokens")
		return
	}

	c.IndentedJSON(http.StatusOK, pair)
}

func (ctrl *AuthController) RefreshToken(c *gin.Context) {

}
