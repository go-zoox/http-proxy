package httpproxy

import (
	"net"
	"net/http"
	"time"

	"github.com/go-zoox/logger"
)

// CONNECT httpbin.org:443 HTTP/1.1
// Host: httpbin.org:443
// Proxy-Connection: keep-alive
// User-Agent: curl/7.84.0
func (hp *httpproxy) handleHTTPs(w http.ResponseWriter, r *http.Request) {
	hp.handleTunneling(w, r)
}

func (hp *httpproxy) handleTunneling(w http.ResponseWriter, r *http.Request) {
	hijacker, ok := w.(http.Hijacker)
	if !ok {
		logger.Errorf("hijack is unsupported (host: %s)", r.Host)
		return
	}

	source, _, err := hijacker.Hijack()
	if err != nil {
		return
	}

	target, err := net.DialTimeout("tcp", r.Host, 60*time.Second)
	if err != nil {
		logger.Errorf("failed to dial to host(%s): %s", r.Host, err)
		return
	}

	source.Write([]byte("HTTP/1.0 200 Connection Established\r\n\r\n"))

	go forward(source, target)
	go forward(target, source)
}
