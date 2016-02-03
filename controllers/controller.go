package controllers

type Controller interface {
	RegisterRoutes(s Server)
}
