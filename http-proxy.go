package httpproxy

import (
	"net/http"
)

type HTTPProxy interface {
	ServeHTTP(w http.ResponseWriter, req *http.Request)
	Run() error
}

func New(opts ...func(cfg *Config)) HTTPProxy {
	cfg := &Config{
		Port: 8080,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	return &httpproxy{
		cfg: cfg,
	}
}

type httpproxy struct {
	cfg *Config
}
