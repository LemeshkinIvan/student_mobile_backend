package auth_api

import (
	auth_models "app/internal/domain/models/auth"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var accessSecret = []byte("your-access-secret")
var refreshSecret = []byte("your-refresh-secret")

func GenerateTokens(userCode string) (auth_models.TokensPairResponse, error) {
	emptyPair := auth_models.TokensPairResponse{Id: -1, AccessToken: "", RefreshToken: ""}

	// Access token
	accessClaims := jwt.MapClaims{
		"user_code": userCode,
		"exp":       time.Now().Add(15 * time.Minute).Unix(),
	}
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(accessSecret)
	if err != nil {
		return emptyPair, err
	}

	// Refresh token
	refreshClaims := jwt.MapClaims{
		"userCode": userCode,
		"exp":      time.Now().Add(2 * time.Hour).Unix(),
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(refreshSecret)
	if err != nil {
		return emptyPair, err
	}

	pair := auth_models.TokensPairResponse{
		Id:           1,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return pair, nil
}
