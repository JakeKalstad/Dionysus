package main

import (
	"github.com/JakeKalstad/Dionysus/config"
	"github.com/JakeKalstad/Dionysus/controllers"
	"github.com/JakeKalstad/Dionysus/server"
	"github.com/gin-gonic/gin"
)

type Context struct {
	*gin.Context
}

func ping(c Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	config := config.Configuration{}
	config.readConfig()
	server := server.Server{
		gin.Default(),
		config,
		[]controllers.Controller{},
	}

	server.Get("/ping", ping)
	server.RegisterControllers()
	server.Run()
}
