package server

import (
	"fmt"
	"github.com/darmiel/yaxc/internal/common"
	"github.com/gofiber/fiber/v2"
	"github.com/muesli/termenv"
	"strings"
)

func (s *yAxCServer) handleGetAnywhere(ctx *fiber.Ctx) (err error) {
	path := strings.TrimSpace(ctx.Params("anywhere"))

	// validate path
	if !common.ValidateAnywherePath(path) {
		return ctx.Status(400).SendString("ERROR: Invalid path")
	}

	var res string
	// get VALUE
	if res, err = s.Backend.GetValue(path); err != nil {
		return
	}

	// Encryption
	if q := ctx.Query("secret"); q != "" {
		if !s.EnableEncryption {
			return errEncryptionNotEnabled
		}
		// do not fail on error
		if encrypt, err := common.Decrypt(res, q); err == nil {
			res = string(encrypt)
		}
	}

	fmt.Println(common.StyleServe(),
		termenv.String(ctx.IP()).Foreground(common.Profile().Color("#DBAB79")),
		"requested",
		termenv.String("value").Foreground(common.Profile().Color("#A8CC8C")),
		termenv.String(path).Foreground(common.Profile().Color("#D290E4")))

	if res == "" {
		ctx.Status(404)
	} else {
		ctx.Status(200)
	}
	return ctx.SendString(res)
}

func (s *yAxCServer) handleGetHashAnywhere(ctx *fiber.Ctx) (err error) {
	path := strings.TrimSpace(ctx.Params("anywhere"))
	var res string
	if res, err = s.Backend.GetHash(path); err != nil {
		return
	}

	fmt.Println(common.StyleServe(),
		termenv.String(ctx.IP()).Foreground(common.Profile().Color("#DBAB79")),
		"requested",
		termenv.String("hash").Foreground(common.Profile().Color("#E88388")),
		termenv.String(path).Foreground(common.Profile().Color("#D290E4")))

	if res == "" {
		ctx.Status(404)
	} else {
		ctx.Status(200)
	}
	return ctx.SendString(res)
}
