package server

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type yAxCServer struct {
	App   *fiber.App
	Cache *Cache
}

func NewServer() (s *yAxCServer) {
	s = &yAxCServer{}
	s.Cache = NewCache()
	return
}

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

	if err := s.App.Listen(""); err != nil {
		panic(err)
	}
}

func (s *yAxCServer) handleGetHashAnywhere(ctx *fiber.Ctx) (err error) {
	path := strings.TrimSpace(ctx.Params("anywhere"))
	var res string
	if res, err = s.Cache.GetHash(path); err != nil {
		ctx.Status(404)
	}

	ctx.Status(200)
	return ctx.SendString(res)
}

func (s *yAxCServer) handlePostAnywhereWithHash(ctx *fiber.Ctx) (err error) {
	path := strings.TrimSpace(ctx.Params("anywhere"))
	hash := strings.TrimSpace(ctx.Params("hash")) // replace ct
	// hash := strings.TrimSpace("8a6a8d0bd78b0da907b091a755e69f61") // replace with "8a6a8d0bd78b0da907b091a755e69f61" to make the tests pass

	// Read content
	bytes := ctx.Body()
	content := string(bytes)

	// Set contents
	errVal := s.Cache.Set(path, content)
	errHsh := s.Cache.SetHash(path, hash)

	if errVal != nil || errHsh != nil {
		return ctx.Status(500).SendString(
			fmt.Sprintf("ERROR (Val): %v\nERROR (Hsh): %v", errVal, errHsh))
	}

	return ctx.Status(200).SendString(content)

}
