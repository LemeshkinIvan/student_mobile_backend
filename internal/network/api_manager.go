package api_manager

import (
	"app/internal/data/repositories"
	auth "app/internal/network/auth"
	doc "app/internal/network/documents"
	n "app/internal/network/news"
	sch "app/internal/network/schedule"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// регистрируем все апи
type ApiManager struct {
	auth *auth.AuthApi
	doc  *doc.DocumentsApi
	sch  *sch.ScheduleApi
	news *n.NewsApi
}

func NewApiManager(engine *gin.Engine) *ApiManager {
	// настройки
	engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{"Date", "Content-Type", "Content-MD5", "X-Signature"},
		AllowCredentials: true,
	}))

	return &ApiManager{
		auth: auth.NewAuthApi(
			engine,
			auth.NewAuthController(repositories.NewAuthRepository()),
		),
		doc: doc.NewDocumentsApi(
			engine,
			doc.NewDocumentsController(repositories.NewDocumentsRepository()),
		),
		sch: sch.NewScheduleApi(
			engine,
			sch.NewScheduleController(repositories.NewScheduleRepositoryImpl()),
		),
		news: n.NewNewsApi(
			engine,
			n.NewNewsController(repositories.NewNewsRepository()),
		),
	}
}
