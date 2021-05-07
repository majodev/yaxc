package server

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
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

	// PROBLEMATIC reuse of ctx buffer: path and hash are stored in cache later.
	// see https://docs.gofiber.io/#zero-allocation
	// path := strings.TrimSpace(ctx.Params("anywhere"))
	// hash := strings.TrimSpace(ctx.Params("hash")) // replace ct

	// fix for https://stackoverflow.com/questions/66930097/race-with-mutex-corrupt-data-in-map
	path := strings.TrimSpace(utils.ImmutableString(ctx.Params("anywhere")))
	hash := strings.TrimSpace(utils.ImmutableString(ctx.Params("hash")))

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
