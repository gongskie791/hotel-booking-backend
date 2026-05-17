package application

import (
	"booking-system/lvl1/internal"
	"booking-system/lvl1/internal/config"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	DB     *config.DB
}

func New() *App {
	return &App{
		Router: gin.Default(),
		DB:     config.NewDB(),
	}
}

func (a *App) Run(addr string) error {
	internal.NewRouter(a.Router, a.DB)
	return a.Router.Run(addr)
}
