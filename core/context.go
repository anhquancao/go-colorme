package core

import (
	"sync"
	"github.com/gin-gonic/gin"
	"colorme.vn/core/registry"
)

type Context struct {
	Server          *gin.Engine
	RegistryManager *registry.RegistryManager
}

var instance *Context
var once sync.Once

func GetContext() *Context {
	once.Do(func() {
		instance = NewContext()
	})
	return instance
}

func NewContext() *Context {
	context := &Context{
		Server:          gin.Default(),
		RegistryManager: registry.NewRegistryManager(),

	}
	return context
}
