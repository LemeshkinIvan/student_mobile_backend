package documents_controller

import (
	documents_repository "app/internal/data/repository/documents"
	service "app/internal/services/document_service"

	"github.com/gin-gonic/gin"
)

type DocumentsController struct {
	DocumentsService *service.DocumentsService
	Repository       *documents_repository.DocumentsRepositoryImpl
}

func NewDocumentsController(
	documentsService *service.DocumentsService,
	repository *documents_repository.DocumentsRepositoryImpl,
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
