package auth_models

type UserResponse struct {
	Uid      string
	FullName string
	Group    Group
}

type UserRequest struct {
	Code string
}
