package repositories

import doc_mod "app/internal/domain/models/documents"

type DocumentsRepositoryImpl struct{}

func NewDocumentsRepository() *DocumentsRepositoryImpl {
	return &DocumentsRepositoryImpl{}
}

func (ur *DocumentsRepositoryImpl) GetDocumentByUid(token string) doc_mod.DocumentResponse {
	mok_user := doc_mod.DocumentResponse{Id: "uid", Url: ""}

	// add error returning type
	return mok_user
}

func (ur *DocumentsRepositoryImpl) GetUserDocuments(code string) []doc_mod.DocumentResponse {
	mok_data := []doc_mod.DocumentResponse{
		{Id: "uid", Url: "ee"},
		{Id: "uid1", Url: "gg"},
	}

	return mok_data
}

func (ur *DocumentsRepositoryImpl) GetAllDocuments(code string) []doc_mod.DocumentResponse {
	mok_data := []doc_mod.DocumentResponse{
		{Id: "uid", Url: "ee"},
		{Id: "uid1", Url: "gg"},
	}

	return mok_data
}
