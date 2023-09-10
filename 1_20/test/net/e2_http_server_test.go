package test

import (
	"golearn/net"
	"testing"
)

// go test -v e2_http_server_test.go
func TestHttpServer(t *testing.T) {
	net.HttpServer()
}
