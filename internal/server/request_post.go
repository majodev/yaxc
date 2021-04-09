package server

import (
	"errors"
	"fmt"
	"github.com/darmiel/yaxc/internal/common"
	"github.com/gofiber/fiber/v2"
	"github.com/muesli/termenv"
	"strings"
	"time"
)

var errEncryptionNotEnabled = errors.New("encryption not enabled")

func (s *yAxCServer) handlePostAnywhere(ctx *fiber.Ctx) (err error) {
	path := strings.TrimSpace(ctx.Params("anywhere"))
	return s.setAnywhereWithHash(ctx, path, "")
}

func (s *yAxCServer) handlePostAnywhereWithHash(ctx *fiber.Ctx) (err error) {
	path := strings.TrimSpace(ctx.Params("anywhere"))
	hash := strings.TrimSpace(ctx.Params("hash"))
	// validate hash
	if !common.ValidateHex(hash) {
		return ctx.Status(400).SendString("ERROR: Invalid hash")
	}
	return s.setAnywhereWithHash(ctx, path, hash)
}

func (s *yAxCServer) setAnywhereWithHash(ctx *fiber.Ctx, path, hash string) (err error) {
	// validate path
	if !common.ValidateAnywherePath(path) {
		return ctx.Status(400).SendString("ERROR: Invalid path")
	}

	// Read content
	bytes := ctx.Body()
	if s.MaxBodyLength > 0 && len(bytes) > s.MaxBodyLength {
		return s.errBodyLen
	}
	content := string(bytes)

	// TTL
	ttl := s.DefaultTTL
	if q := ctx.Query("ttl"); q != "" {
		if ttl, err = time.ParseDuration(q); err != nil {
			return
		}
	}

	// Encryption
	if q := ctx.Query("secret"); q != "" {
		if !s.EnableEncryption {
			return errEncryptionNotEnabled
		}
		// fail on error
		encrypt, err := common.Encrypt(content, q)
		if err != nil {
			return err
		}
		content = string(encrypt)
	}

	// Check if ttl is valid
	if (s.MinTTL != 0 && s.MinTTL > ttl) || (s.MaxTTL != 0 && s.MaxTTL < ttl) {
		return ctx.Status(400).SendString("ERROR: TTL out of range")
	}

	// generate hash
	if hash == "" {
		hash = common.MD5Hash(content)
	}

	// Set contents
	if err := s.Backend.Set(path, content, hash, ttl); err != nil {
		log.Warning("ERROR saving anywhere:", err)
		return ctx.Status(500).SendString(fmt.Sprintf("ERROR: %v", err))
	}

	fmt.Println(common.StyleServe(),
		termenv.String(ctx.IP()).Foreground(common.Profile().Color("#DBAB79")),
		"updated",
		termenv.String(path).Foreground(common.Profile().Color("#D290E4")),
		"with hash",
		termenv.String(hash).Foreground(common.Profile().Color("#71BEF2")))

	return ctx.Status(200).SendString(content)
}
