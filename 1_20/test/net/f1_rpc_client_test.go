package test

import (
	"golearn/net"
	"testing"
)

// go test -v f1_rpc_client_test.go
func TestRpcClient(t *testing.T) {
	net.RpcClient()
}
