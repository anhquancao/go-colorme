package app

import (
	"colorme.vn/core"
	"colorme.vn/core/service"
	"colorme.vn/registry"
	"colorme.vn/router"
)

type App struct {
	service *service.Service
	context *core.Context
}

func NewApp() *App {
	app := &App{
		context: core.GetContext(),
		service: service.NewService(),
	}
	return app
}

func (app *App) Init() {
	app.context.RegistryManager.RegisterControllerRegistry(registry.GetControllerRegistry())
	router.RegisterBaseRouter(app.context)

}

func (app *App) Run() {
	server := app.context.Server
	server.Run("127.0.0.1:8080")
}
