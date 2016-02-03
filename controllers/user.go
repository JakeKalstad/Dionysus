package controllers

import "github.com/gin-gonic/gin"

type UserController struct {
}

func (c UserController) RegisterRoutes(s Server) {
	s.Post("/user", c.Save)
	s.Get("/user", c.Retrieve)
	s.Delete("/user", c.Delete)
}

func (c *UserController) Save(ctx *gin.Context) {

}

func (c *UserController) Delete(ctx *gin.Context) {

}

func (c *UserController) Retrieve(ctx *gin.Context) {

}
