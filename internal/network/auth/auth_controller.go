package auth_api

import (
	r "app/internal/data/repositories"
	models "app/internal/domain/models/auth"
	handler "app/internal/network/handlers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	Repository *r.AuthRepositoryImpl
}

func NewAuthController(
	repository *r.AuthRepositoryImpl,
) *AuthController {
	return &AuthController{
		Repository: repository,
	}
}

func (ctrl *AuthController) GetUsers(c *gin.Context) {
	var request models.UserRequest

	if err := c.BindJSON(&request); err != nil {
		var alert string = "bad client getUser request " + err.Error()
		log.Panic(alert)

		c.JSON(http.StatusBadRequest, handler.ErrorRessponse{
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

		c.JSON(http.StatusBadRequest, handler.ErrorRessponse{
			Message: "byg",
		})

		return
	}

	if tokenRequest.Code != "12" {
		var alert string = "user not exist"
		c.JSON(http.StatusBadRequest, handler.ErrorRessponse{
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

		c.JSON(http.StatusBadRequest, handler.ErrorRessponse{
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
