package server

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (s *yAxCServer) handleGetHashAnywhere(ctx *fiber.Ctx) (err error) {
	path := strings.TrimSpace(ctx.Params("anywhere"))
	var res string
	if res, err = s.Backend.GetHash(path); err != nil {
		ctx.Status(404)
	}

	ctx.Status(200)
	return ctx.SendString(res)
}
