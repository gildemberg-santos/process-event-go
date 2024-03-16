package router

import (
	"github.com/gildemberg-santos/process-event-go/internal/controller"
	"github.com/gin-gonic/gin"
)

type Router struct {
	Route *gin.Engine
}

func NewRoute() *Router {
	return &Router{
		Route: gin.Default(),
	}
}

func (r *Router) setting() {
	r.Route.Use(gin.Recovery())
	r.Route.Use(controller.AuthMiddleware)
}

func (r *Router) Run() {
	r.Route.POST("/login", controller.Auth)
	r.setting()
	r.Route.GET("/event", controller.EventController)
	r.Route.Run()
}
