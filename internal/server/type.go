package server

import (
	"errors"
	"os"
	"time"

	"github.com/darmiel/yaxc/internal/bcache"
	"github.com/gofiber/fiber/v2"
)

type YAxCConfig struct {
	PrefixVal string

	// Address
	BindAddress string // required
	// Timeout
	DefaultTTL time.Duration // 0 -> infinite
	MinTTL     time.Duration // == MaxTTL -> cannot specify TTL
	MaxTTL     time.Duration // == MinTTL -> cannot specify TTL
	// Other
	MaxBodyLength    int
	EnableEncryption bool
	ProxyHeader      string
}

type yAxCServer struct {
	*YAxCConfig
	App        *fiber.App
	Backend    Backend
	errBodyLen error
}

func NewServer(cfg *YAxCConfig) (s *yAxCServer) {
	s = &yAxCServer{
		YAxCConfig: cfg,
		errBodyLen: errors.New("exceeded max body length"),
	}

	s.Backend = &CacheBackend{
		cache:   bcache.NewCache(),
		errCast: errors.New("not a string"),
	}

	if s.Backend == nil {
		os.Exit(1)
		return
	}

	return
}
