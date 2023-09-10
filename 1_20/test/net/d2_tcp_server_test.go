package test

import (
	"golearn/net"
	"testing"
)

// go test -v d2_tcp_server_test.go
func TestTcpServer(t *testing.T) {
	net.TcpServer()
}
