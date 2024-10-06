# HTTP Proxy - Light HTTP Proxy Server

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/http-proxy)](https://pkg.go.dev/github.com/go-zoox/http-proxy)
[![Build Status](https://github.com/go-zoox/http-proxy/actions/workflows/lint.yml/badge.svg?branch=master)](https://github.com/go-zoox/http-proxy/actions/workflows/lint.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/http-proxy)](https://goreportcard.com/report/github.com/go-zoox/http-proxy)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/http-proxy/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/http-proxy?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/http-proxy.svg)](https://github.com/go-zoox/http-proxy/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/http-proxy.svg?label=Release)](https://github.com/go-zoox/http-proxy/tags)


## Installation
To install the package, run:
```bash
go get -u github.com/go-zoox/http-proxy
```

## Quick Start

```go
package main

import (
	"fmt"
	"net/http"

	"github.com/go-zoox/http-proxy"
)

func main() {
	fmt.Println("Starting proxy at http://127.0.0.1:1080 ...")

	http.ListenAndServe(":1080", httpproxy.New())
}

// curl --proxy http://127.0.0.1:1080 http://httpbin.org
```

## Inspiration
* https://medium.com/@mlowicki/http-s-proxy-in-golang-in-less-than-100-lines-of-code-6a51c2f2c38c

## License
GoZoox is released under the [MIT License](./LICENSE).
