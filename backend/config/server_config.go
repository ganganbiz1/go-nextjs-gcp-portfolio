package config

import (
	"os"
	"strconv"

	"github.com/ganganbiz1/go-nextjs-gcp-portfolio/backend/domain"
)

type ServerConfig struct {
	Port int
}

func NewServerConfig() (*ServerConfig, error) {
	p, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		return nil, handleError(domain.ErrTypeConvert, err)
	}
	return &ServerConfig{
		Port: p,
	}, nil
}
