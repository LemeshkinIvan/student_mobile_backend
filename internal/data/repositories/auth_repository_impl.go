package repositories

import (
	"app/internal/bootstrap/db"
	auth_models "app/internal/domain/models/auth"
	"errors"
)

type AuthRepositoryImpl struct {
	db *db.Postgres
}

func NewAuthRepository(db *db.Postgres) *AuthRepositoryImpl {
	return &AuthRepositoryImpl{
		db: db,
	}
}

func (ur *AuthRepositoryImpl) GetUserByAccessToken(token string) auth_models.UserResponse {
	mok_group := auth_models.Group{Id: token, Name: "group1"}
	mok_user := auth_models.UserResponse{Uid: token, FullName: "user1", Group: mok_group}

	// add error returning type
	return mok_user
}

func (ur *AuthRepositoryImpl) GetTokenPair(code string) (auth_models.TokensPairResponse, error) {
	mok_tokens := auth_models.TokensPairResponse{}
	e := errors.New("")
	return mok_tokens, e
}

func (ur *AuthRepositoryImpl) RefreshToken(code string) (auth_models.TokensPairResponse, error) {
	mok_tokens := auth_models.TokensPairResponse{}
	e := errors.New("")
	return mok_tokens, e
}
