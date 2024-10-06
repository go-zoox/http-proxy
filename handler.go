package httpproxy

import (
	"net/http"

	"github.com/go-zoox/logger"
)

func (hp *httpproxy) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	// proxy authorization
	proxyAuthorization := req.Header.Get("Proxy-Authorization")
	if hp.cfg.Username != "" || hp.cfg.Password != "" {
		if u, p, ok := parseBasicAuth(proxyAuthorization); !ok {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		} else if u != hp.cfg.Username || p != hp.cfg.Password {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
	}

	// HTTPS
	if req.Method == http.MethodConnect {
		logger.Infof("[https][%s][%s]", req.RemoteAddr, req.Host)

		hp.handleHTTPs(rw, req)
	} else {
		// HTTP
		logger.Infof("[http][%s][%s]", req.RemoteAddr, req.Host)

		hp.handleHTTP(rw, req)
	}
}
