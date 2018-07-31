package router

import (
	"colorme.vn/core"
	"colorme.vn/controller/base"
)

func RegisterBaseRouter(context *core.Context) {
	server := context.Server
	server.GET("base", base.BaseGraphQL)
}
