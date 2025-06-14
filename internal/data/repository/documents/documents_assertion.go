package documents_repository

import (
	documents_repository_interface "app/internal/domain/repositories/documents"
)

var _ documents_repository_interface.DocumentsRepository = (*DocumentsRepositoryImpl)(nil)
