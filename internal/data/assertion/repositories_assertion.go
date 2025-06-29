package repository_assertion

import (
	impl "app/internal/data/repositories"

	auth "app/internal/domain/repositories/auth"
	doc "app/internal/domain/repositories/documents"
	news "app/internal/domain/repositories/news"
	sch "app/internal/domain/repositories/schedule"
)

// явное определение контракта интерфейсов
var _ auth.AuthRepository = (*impl.AuthRepositoryImpl)(nil)

var _ doc.DocumentsRepository = (*impl.DocumentsRepositoryImpl)(nil)

var _ sch.SheduleRepository = (*impl.ScheduleRepositoryImpl)(nil)

var _ news.NewsRepositoryImpl = (*impl.NewsRepositoryImpl)(nil)
