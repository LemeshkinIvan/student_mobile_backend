package doc_api

import (
	r "app/internal/data/repositories"

	"github.com/gin-gonic/gin"
)

type DocumentsController struct {
	Repository *r.DocumentsRepositoryImpl
}

func NewDocumentsController(
	repository *r.DocumentsRepositoryImpl,
) *DocumentsController {
	return &DocumentsController{
		Repository: repository,
	}
}

func (ctrl *DocumentsController) GetDocumentByUid(c *gin.Context) {

}

func (ctrl *DocumentsController) GetUserDocuments(c *gin.Context) {

}

func (ctrl *DocumentsController) GetAllDocuments(c *gin.Context) {

}
