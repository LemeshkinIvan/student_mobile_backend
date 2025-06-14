package documents_repository_interface

import documents_models "app/internal/domain/models/documents"

type DocumentsRepository interface {
	GetDocumentByUid(token string) documents_models.DocumentResponse

	GetUserDocuments(code string) []documents_models.DocumentResponse

	GetAllDocuments(code string) []documents_models.DocumentResponse
}
