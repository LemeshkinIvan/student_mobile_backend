package auth_models

type TokensPairResponse struct {
	Id           int32
	AccessToken  string
	RefreshToken string
}

type TokensPairRequest struct {
	Code string
}

type TokensRefreshRequest struct {
	AccessToken  string
	RefreshToken string
}
