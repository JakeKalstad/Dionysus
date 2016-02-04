package controllers

type UserController struct {
}

func (c UserController) RegisterRoutes(s Server) {
	s.Post("/user", c.Save)
	s.Get("/user", c.Retrieve)
	s.Delete("/user", c.Delete)
}

func (c *UserController) Save(ctx Context) {

}

func (c *UserController) Delete(ctx Context) {

}

func (c *UserController) Retrieve(ctx Context) {

}
