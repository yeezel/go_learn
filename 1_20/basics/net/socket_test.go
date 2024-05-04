package net

import (
	"fmt"
	"io"
	"net"
	"os"
	"testing"
)

func TestSocket(t *testing.T) {
	SocketDemo()
}
func SocketDemo() {
	var (
		host          = "www.apache.org"
		port          = "80"
		remote        = host + ":" + port
		msg    string = "GET / \n"
		data          = make([]uint8, 4096)
		read          = true
		count         = 0
	)

	//连接
	// conn, err := net.Dial("tcp", "192.0.32.10:80") // tcp ipv4
	// checkConnection(conn, err)
	// conn, err = net.Dial("udp", "192.0.32.10:80") // udp
	// checkConnection(conn, err)
	// conn, err = net.Dial("tcp", "[2620:0:2d0:200::10]:80") // tcp ipv6
	// checkConnection(conn, err)

	// 创建一个 socket
	con, err := net.Dial("tcp", remote)
	checkConnection(con, err)
	// 发送我们的消息，一个 http GET 请求
	io.WriteString(con, msg)
	// 读取服务器的响应
	for read {
		count, err = con.Read(data)
		read = (err == nil)
		fmt.Printf(string(data[0:count]))
	}
	con.Close()
}

func checkConnection(conn net.Conn, err error) {
	if err != nil {
		fmt.Printf("error %v connecting!", err)
		os.Exit(1)
	}
	fmt.Printf("Connection is made with %v\n", conn)
}
