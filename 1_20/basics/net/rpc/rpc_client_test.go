package net

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/rpc"
	"os"
	"testing"
)

// go test -v f1_rpc_client_test.go
func TestRpcHttpClient(t *testing.T) {
	RpcHttpClient()
}

const serverAddress = "localhost"

func RpcHttpClient() {
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("Error dialing:", err)
	}
	args := &Args{7, 8}
	var reply int
	// 同步调用
	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Args error:", err)
	}
	fmt.Printf("Args: %d * %d = %d", args.N, args.M, reply)

	// 异步调用
	call1 := client.Go("Args.Multiply", args, &reply, nil) //最后一个参数值为 nil ，调用完成后会创建一个新的通道
	// replyCall := <-call1.Done
	<-call1.Done
}

func RpcSSLClient() {
	// ssl rpc
	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, _ := tls.Dial("tcp", serverAddress+":1234", config)
	defer conn.Close()
	client := rpc.NewClient(conn)

	args := &Args{7, 8}
	var result int
	if err := client.Call("Cal.Square", args, &result); err != nil {
		log.Fatal("Failed to call Cal.Square. ", err)
	}

	log.Printf("%d^2 = %d", args.N, args.M)
}

// 单向认证, 对应RpcSSLServer
func RpcSSLClient1() {
	// 添加客户端对服务器端鉴权
	certPool := x509.NewCertPool()
	certBytes, err := ioutil.ReadFile("../server/server.crt")
	if err != nil {
		log.Fatal("Failed to read server.crt")
	}
	certPool.AppendCertsFromPEM(certBytes)

	// ssl rpc
	config := &tls.Config{
		RootCAs: certPool,
	}

	conn, _ := tls.Dial("tcp", serverAddress+":1234", config)
	defer conn.Close()
	client := rpc.NewClient(conn)

	args := &Args{7, 8}
	var result int
	if err := client.Call("Cal.Square", args, &result); err != nil {
		log.Fatal("Failed to call Cal.Square. ", err)
	}

	log.Printf("%d^2 = %d", args.N, args.M)
}

// 双向认证
func RpcSSLAuthClient() {
	cert, _ := tls.LoadX509KeyPair("client.crt", "client.key")
	certPool := x509.NewCertPool()
	certBytes, _ := os.ReadFile("../server/server.crt")
	certPool.AppendCertsFromPEM(certBytes)
	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
	}

	conn, _ := tls.Dial("tcp", serverAddress+":1234", config)
	defer conn.Close()
	client := rpc.NewClient(conn)

	args := &Args{7, 8}
	var result int
	if err := client.Call("Cal.Square", args, &result); err != nil {
		log.Fatal("Failed to call Cal.Square. ", err)
	}

	log.Printf("%d^2 = %d", args.N, args.M)
}
