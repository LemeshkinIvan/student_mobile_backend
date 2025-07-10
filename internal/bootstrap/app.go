package bootstrap

import (
	"app/internal/bootstrap/db"
	"app/internal/bootstrap/env"
	l "app/internal/common/logger"
	auth_api "app/internal/network/auth"
	doc_api "app/internal/network/documents"
	news_api "app/internal/network/news"
	schedule_api "app/internal/network/schedule"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type App struct {
	Database *db.Postgres
	Config   *env.EnvConfig

	// network
	Engine *gin.Engine
	auth   *auth_api.AuthApi
	doc    *doc_api.DocumentsApi
	sch    *schedule_api.ScheduleApi
	news   *news_api.NewsApi
}

func InitApp() *App {
	conf, err := env.MustLoadEnv()
	if err != nil {
		l.Logg.Error(err.Error())
	}
	l.Logg.Info("app conf was init")

	postgres, err := db.Connect(
		conf.POSTGRES_PASSWORD,
		conf.POSTGRES_USER,
		conf.POSTGRES_PORT,
		conf.POSTGRES_URL,
		conf.DATABASE_NAME,
	)

	if err != nil {
		l.Logg.Error(err.Error())
	}
	l.Logg.Info("database was init")

	// инициализируем движок рутов
	g := gin.Default()

	// https://habr.com/ru/companies/macloud/articles/553826/
	g.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	}))

	auth_api := auth_api.NewAuthApi(g, postgres)
	l.Logg.Info("auth api was init")

	doc_api := doc_api.NewDocumentsApi(g, postgres)
	l.Logg.Info("doc_api was init")

	schedule_api := schedule_api.NewScheduleApi(g, postgres)
	l.Logg.Info("schedule_api was init")

	news_api := news_api.NewNewsApi(g, postgres)
	l.Logg.Info("news_api was init")

	return &App{
		Config:   conf,
		Database: postgres,
		Engine:   g,

		// api
		auth: auth_api,
		doc:  doc_api,
		news: news_api,
		sch:  schedule_api,
	}
}

func InitAppProd() *App {
	return nil
}

// отключение всего и вся
func (app *App) DisposeApp() {
	app.Database.Disconnect()
}
