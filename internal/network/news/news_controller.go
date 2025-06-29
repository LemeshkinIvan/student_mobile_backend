package news_api

import (
	"app/internal/data/repositories"

	"github.com/gin-gonic/gin"
)

type NewsController struct {
	Repository *repositories.NewsRepositoryImpl
}

func NewNewsController(r *repositories.NewsRepositoryImpl) *NewsController {
	return &NewsController{
		Repository: r,
	}
}

func (n *NewsController) GetNews(c *gin.Context) {

}
