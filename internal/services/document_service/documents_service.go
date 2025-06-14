package documents_service

type DocumentsService struct {
	// можно подключать репозитории
	// сервис - это сборище useCase
}

func NewDocumentsService() *DocumentsService {
	return &DocumentsService{}
}

func GetAllDocuments() *DocumentsService {
	return &DocumentsService{}
}

func (s *DocumentsService) GetDocumentsByUid() []string {
	return []string{}
}

func (s *DocumentsService) GetUserDocuments() []string {
	return []string{}
}
