package auth_repository_interface

import auth_models "app/internal/domain/models/auth"

type AuthRepository interface {
	GetUserByAccessToken(token string) auth_models.UserResponse

	GetTokenPair(code string) (auth_models.TokensPairResponse, error)

	RefreshToken(code string) (auth_models.TokensPairResponse, error)
}
