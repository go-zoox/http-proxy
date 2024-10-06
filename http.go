package httpproxy

import (
	"io"
	"net/http"
)

// GET http://httpbin.org/ HTTP/1.1
// Host: httpbin.org
// Proxy-Connection: keep-alive
// User-Agent: curl/7.84.0
func (hp *httpproxy) handleHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	transport := http.DefaultTransport
	outReq := r.Clone(ctx)

	outRes, err := transport.RoundTrip(outReq)
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		w.Write([]byte(err.Error()))
		return
	}
	defer outRes.Body.Close()

	copyHeader(w.Header(), outRes.Header)

	w.WriteHeader(outRes.StatusCode)

	io.Copy(w, outRes.Body)
}
