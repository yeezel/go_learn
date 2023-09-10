package test

import (
	"golearn/net"
	"testing"
)

// go test -v e1_http_client_test.go
func TestHttpClient(t *testing.T) {
	net.HttpClient()
}
