package main

import (
	"fmt"
	"log"
	"net/rpc"
	"./rpc_object"
)

const serverAddress = "localhost"

func main() {
	client, err := rpc.DialHTTP("tcp", serverAddress+":1234")
	if err != nil {
		log.Fatal("Error dialing:", err)
	}
	// Synchronous call
	args := &a240_rpc_objects.Args{7, 8}
	var reply int
	err = client.Call("Args.Multiply", args, &reply)
	if err != nil {
		log.Fatal("Args error:", err)
	}
	fmt.Printf("Args: %d * %d = %d", args.N, args.M, reply)

	// 这是异步方式
	call1 := client.Go("Args.Multiply", args, &reply, nil) //最后一个参数值为 nil ，调用完成后会创建一个新的通道
	replyCall := <-call1.Done
}
