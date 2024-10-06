package httpproxy

import (
	"fmt"
	"net/http"

	"github.com/go-zoox/logger"
)

func (hp *httpproxy) Run() error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", hp.cfg.Port),
		Handler: hp,
	}

	logger.Infof("server started at: 0.0.0.0:%d", hp.cfg.Port)

	return server.ListenAndServe()
}
