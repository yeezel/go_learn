package main

import (
	"errors"
	"fmt"
	"time"
)

func main() {
	//返回一个通道而不关闭
	tick := time.Tick(1e8)  //以1e8纳秒为周期返回一个通道
	boom := time.After(5e8) //在5e8纳秒后返回一个通道
BOOMS:
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			break BOOMS
		default:
			fmt.Println("    .")
			time.Sleep(5e7)
		}
	}

	//场景一：超时
	timeout()

	//取消耗时很长的同步调用
	chanel()

}

func chanel() {
	// 缓冲大小设置为 1 是必要的，可以避免协程死锁以及确保超时的通道可以被垃圾回收
	ch := make(chan error, 1)
	go func() { ch <- errors.New("long time") }()
	select {
	case resp := <-ch:
		println(resp)
	case <-time.After(8e9):
		// call timed out
		break
	}
}

func timeout() {
	datas := make(chan string)
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9) // one second
		timeout <- true
	}()
	select {
	case v := <-datas:
		println(v)
	case <-timeout:
		println("timeout")
		break
	}
}
