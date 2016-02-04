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
