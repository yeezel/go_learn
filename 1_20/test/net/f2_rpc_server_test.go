package test

import (
	"golearn/net"
	"testing"
)

// go test -v f2_rpc_server_test.go
func TestRpcServer(t *testing.T) {
	net.RpcServer()
}
