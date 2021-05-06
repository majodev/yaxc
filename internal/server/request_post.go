package server

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (s *yAxCServer) handlePostAnywhereWithHash(ctx *fiber.Ctx) (err error) {
	path := strings.TrimSpace(ctx.Params("anywhere"))
	hash := strings.TrimSpace(ctx.Params("hash"))

	// Read content
	bytes := ctx.Body()
	content := string(bytes)

	// Set contents
	errVal := s.Backend.Set(path, content)
	errHsh := s.Backend.SetHash(path, hash)

	if errVal != nil || errHsh != nil {
		return ctx.Status(500).SendString(
			fmt.Sprintf("ERROR (Val): %v\nERROR (Hsh): %v", errVal, errHsh))
	}

	return ctx.Status(200).SendString(content)

}
