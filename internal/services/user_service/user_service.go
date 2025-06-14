package services

type UserService struct {
	// можно подключать репозитории
	// сервис - это сборище useCase
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) GetAllUsers() []string {
	return []string{}
}
