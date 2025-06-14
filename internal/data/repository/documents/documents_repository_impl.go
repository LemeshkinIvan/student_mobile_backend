package documents_repository

import documents_models "app/internal/domain/models/documents"

type DocumentsRepositoryImpl struct{}

func NewDocumentsRepository() *DocumentsRepositoryImpl {
	return &DocumentsRepositoryImpl{}
}

func (ur *DocumentsRepositoryImpl) GetDocumentByUid(token string) documents_models.DocumentResponse {
	mok_user := documents_models.DocumentResponse{Id: "uid", Url: ""}

	// add error returning type
	return mok_user
}

func (ur *DocumentsRepositoryImpl) GetUserDocuments(code string) []documents_models.DocumentResponse {
	mok_data := []documents_models.DocumentResponse{
		documents_models.DocumentResponse{Id: "uid", Url: "ee"},
		documents_models.DocumentResponse{Id: "uid1", Url: "gg"},
	}

	return mok_data
}

func (ur *DocumentsRepositoryImpl) GetAllDocuments(code string) []documents_models.DocumentResponse {
	mok_data := []documents_models.DocumentResponse{
		documents_models.DocumentResponse{Id: "uid", Url: "ee"},
		documents_models.DocumentResponse{Id: "uid1", Url: "gg"},
	}

	return mok_data
}
