package bootstrap

import (
	api_manager "app/internal/network"

	"github.com/gin-gonic/gin"
)

type App struct {
	GinEngine *gin.Engine
	Api       *api_manager.ApiManager
	Db        *DatabaseManager
	Env       *Env
}

func InitApp() *App {
	e := gin.Default()
	api_manager := api_manager.NewApiManager(e)
	database := NewDatabaseManager()
	env := NewEnv()

	return &App{
		Db:        database,
		Env:       env,
		Api:       api_manager,
		GinEngine: e,
	}
}
