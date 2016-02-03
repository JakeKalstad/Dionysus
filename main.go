package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/JakeKalstad/Dionysus/controllers"
	"github.com/gin-gonic/gin"
)

type Configuration struct {
}

type Server struct {
	*gin.Engine
	Config      Configuration
	Controllers []controllers.Controller
}

func (s Server) RegisterControllers() {
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

func (c Configuration) readConfig() {
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		fmt.Println("error:", err)
	}
}

type Context struct {
	*gin.Context
}

func ping(c Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func main() {
	config := Configuration{}
	config.readConfig()
	server := Server{
		gin.Default(),
		config,
		[]controllers.Controller{},
	}

	server.Get("/ping", ping)
	server.RegisterControllers()
	server.Run()
}
