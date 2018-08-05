package app

import (
	"colorme.vn/core"
	"colorme.vn/core/service"
	"colorme.vn/registry"
	"colorme.vn/router"
	"github.com/gin-gonic/gin"
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

func setupGraphQLUI(server *gin.Engine) {
	server.Static("/graphqlui", "./public/graphqlui")
	server.Static("/static", "./public/graphqlui/static")
}

func (app *App) Init() {
	app.context.RegistryManager.RegisterControllerRegistry(registry.GetControllerRegistry())
	router.RegisterGraphQLRouter(app.context)
	server := core.GetContext().Server

	server.Static("/assets", "./public/assets")

	setupGraphQLUI(server)
	//server.LoadHTMLGlob("views/*")

}

func (app *App) Run() {
	server := app.context.Server
	server.Run(":8080")
}
