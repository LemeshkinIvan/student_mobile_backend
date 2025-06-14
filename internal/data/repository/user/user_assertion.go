package user_repository

import (
	user_repository_interface "app/internal/domain/repositories/user"
)

// явное определение контракта интерфейсов
var _ user_repository_interface.UserRepository = (*UserRepositoryImpl)(nil)
