package router

import (
	"time"

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
	r.Route.GET("/", controller.Home)
	r.Route.GET("/sleep", func(ctx *gin.Context) {
		time.Sleep(60 * time.Second)
		ctx.JSON(200, gin.H{"message": "sleep"})
	})
	r.Route.Run()
}
