package user_repository

import (
	auth_models "app/internal/domain/models/auth"
	auth_utils "app/internal/utils"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (ur *UserRepositoryImpl) GetUserByAccessToken(token string) auth_models.UserResponse {
	mok_group := auth_models.Group{Id: token, Name: "group1"}
	mok_user := auth_models.UserResponse{Uid: token, FullName: "user1", Group: mok_group}

	// add error returning type
	return mok_user
}

func (ur *UserRepositoryImpl) GetTokenPair(code string) (auth_models.TokensPairResponse, error) {
	mok_tokens, error := auth_utils.GenerateTokens(code)
	return mok_tokens, error
}

func (ur *UserRepositoryImpl) RefreshToken(code string) (auth_models.TokensPairResponse, error) {
	mok_tokens, error := auth_utils.GenerateTokens(code)
	return mok_tokens, error
}
