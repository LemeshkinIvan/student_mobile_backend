package auth_controllers

import (
	error_resp "app/internal/api/controllers/common_ctrl"
	user_repository "app/internal/data/repository/user"
	models "app/internal/domain/models/auth"
	services "app/internal/services/user_service"
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
	var request models.UserRequest

	if err := c.BindJSON(&request); err != nil {
		var alert string = "bad client getUser request " + err.Error()
		log.Panic(alert)

		c.JSON(http.StatusBadRequest, error_resp.ErrorRessponse{
			Message: err.Error(),
		})
	}

	var user = ctrl.Repository.GetUserByAccessToken(request.Code)

	c.IndentedJSON(http.StatusOK, user)
}

func (ctrl *AuthController) GetTokens(c *gin.Context) {
	var tokenRequest models.TokensPairRequest

	if err := c.BindJSON(&tokenRequest); err != nil {
		var alert string = "bad client token request " + err.Error()
		log.Panic(alert)

		c.JSON(http.StatusBadRequest, error_resp.ErrorRessponse{
			Message: "byg",
		})

		return
	}

	if tokenRequest.Code != "12" {
		var alert string = "user not exist"
		c.JSON(http.StatusBadRequest, error_resp.ErrorRessponse{
			Message: alert,
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
	var tokenRefresh models.TokensRefreshRequest

	if err := c.BindJSON(&tokenRefresh); err != nil {
		var alert string = "bad client token request " + err.Error()
		log.Panic(alert)

		c.JSON(http.StatusBadRequest, error_resp.ErrorRessponse{
			Message: "byg",
		})

		return
	}

	var pair, err = ctrl.Repository.GetTokenPair(tokenRefresh.AccessToken)

	if err != nil {
		var alert string = "cant generate tokens " + err.Error()
		log.Panic(alert)

		c.JSON(http.StatusInternalServerError, "tokens")
		return
	}

	c.IndentedJSON(http.StatusOK, pair)
}
