package server

import (
	"github.com/JakeKalstad/Dionysus/controllers"
	"github.com/JakeKalstad/Dionysus/server"
	"github.com/gin-gonic/gin"
)

type Server struct {
	*gin.Engine
	Config      Configuration
	Controllers []controllers.Controller
}

func (s server.Server) RegisterControllers() {
	for _, c := range s.Controllers {
		c.RegisterRoutes(s)
	}
}

func (s Server) Get(url string, action func(c Context)) {
	s.GET(url, func(ctx *gin.Context) { action(Context{ctx}) })
}

func (s Server) Post(url string, action func(c Context)) {
	s.POST(url, func(ctx *gin.Context) { action(Context{ctx}) })
}

func (s Server) Delete(url string, action func(c Context)) {
	s.DELETE(url, func(ctx *gin.Context) { action(Context{ctx}) })
}

func (s Server) Put(url string, action func(c Context)) {
	s.PUT(url, func(ctx *gin.Context) { action(Context{ctx}) })
}
