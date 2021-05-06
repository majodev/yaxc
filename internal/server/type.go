package server

import (
	"github.com/gofiber/fiber/v2"
)

type YAxCConfig struct {
	// Address
	BindAddress string // required
}

type yAxCServer struct {
	*YAxCConfig
	App        *fiber.App
	Backend    *Cache
	errBodyLen error
}

func NewServer(cfg *YAxCConfig) (s *yAxCServer) {
	s = &yAxCServer{
		YAxCConfig: cfg,
	}

	s.Backend = NewCache()

	return
}
