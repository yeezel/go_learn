package net

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"testing"
	"time"
)

// go test -v f2_rpc_server_test.go
func TestRpcHttpServer(t *testing.T) {
	RpcHttpServer()
}

func RpcHttpServer() {
	calc := new(Args)
	rpc.Register(calc)
	// http rpc
	rpc.HandleHTTP()
	listener, e := net.Listen("tcp", "localhost:1234")
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}
	go http.Serve(listener, nil)
	time.Sleep(1000e9)

}

func RpcSSLServer() {
	calc := new(Args)
	rpc.Register(calc)
	// ssl rpc
	certFile, keyFile := "", ""
	cert, _ := tls.LoadX509KeyPair(certFile, keyFile)
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
	}
	listener, e := tls.Listen("tcp", ":1234", config)
	log.Printf("Serving RPC server on port %d", 1234)
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}
	for {
		conn, _ := listener.Accept()
		defer conn.Close()
		go rpc.ServeConn(conn)
	}
}

// 双向认证
func RpcSSLAuthServer() {
	calc := new(Args)
	rpc.Register(calc)
	// ssl rpc
	certFile, keyFile := "client.crt", "client.key"
	cert, _ := tls.LoadX509KeyPair(certFile, keyFile)
	certPool := x509.NewCertPool()
	certBytes, _ := ioutil.ReadFile("../client/client.crt")
	certPool.AppendCertsFromPEM(certBytes)
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    certPool,
	}
	listener, e := tls.Listen("tcp", ":1234", config)
	log.Printf("Serving RPC server on port %d", 1234)
	if e != nil {
		log.Fatal("Starting RPC-server -listen error:", e)
	}
	for {
		conn, _ := listener.Accept()
		defer conn.Close()
		go rpc.ServeConn(conn)
	}
}
