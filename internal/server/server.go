package server

import (
	"github.com/gofiber/fiber/v2"
)

func (s *yAxCServer) StartInternal() {
	cfg := &fiber.Config{}
	s.App = fiber.New(*cfg)

	// GET hash
	s.App.Get("/hash/:anywhere", s.handleGetHashAnywhere)

	// SET contents, custom hash
	s.App.Post("/:anywhere/:hash", s.handlePostAnywhereWithHash)
}

func (s *yAxCServer) Start() {
	s.StartInternal()

	if err := s.App.Listen(s.BindAddress); err != nil {
		panic(err)
	}
}
