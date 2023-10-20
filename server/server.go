package server

import (
	"anime/database"

	"github.com/gofiber/fiber/v2"
)

type Server struct {
	db     database.Database
	router *fiber.App
}

func (s Server) initHandlers() {
	s.router.Get("/users/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")

		return c.JSON(s.db.GetUser(name))
	})
}

func (s Server) Start(addr string) {
	s.router.Listen(addr)
}

func NewServer(db database.Database) Server {
	s := Server{
		db:     db,
		router: fiber.New(),
	}

	s.initHandlers()

	return s
}
