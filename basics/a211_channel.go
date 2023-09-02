package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建通道，协程在通道内进行数据交换
	// 无缓冲，一次只能包含一个元素
	isok := make(chan string)
	// 带缓冲通道
	buf := 100
	ch := make(chan string, buf)

	go sendData(ch)
	for i := 0; i < 3; i++ {
		go getData(ch, isok)
	}

	// time.Sleep(5e9)
	for i := 0; i < 3; i++ {
		// 使用通道阻塞特性检测是否协程是否处理完成
		<-isok
	}

	sendChan := make(chan int)
	receiveChan := make(chan string)
	//使用 for-range 语句会自动检测通道是否关闭
	go processChannel(sendChan, receiveChan)

}

// 有方向的通道参数
func sendData(ch chan<- string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
	//通道可以被显式的关闭；尽管它们和文件不同：不必每次都关闭。
	//只有在当需要告诉接收者不会再提供新的值的时候，才需要关闭通道。
	//只有发送者需要关闭通道，接收者永远不会需要。
	close(ch)
}

// 有方向的通道参数
func getData(ch <-chan string, isok chan<- string) {
	for {
		input, open := <-ch
		if !open { // 检测通道有没有被阻塞，false为阻塞
			break
		}
		fmt.Printf("%s ", input)
	}

	time.Sleep(1e9)
	println("get data end")
	isok <- "ok"
}

func processChannel(in <-chan int, out chan<- string) {
	for inValue := range in {
		result := string(inValue) /// processing inValue
		out <- result
	}
	//filter
	// for {
	// 	i := <-in // Receive value of new variable 'i' from 'in'.
	// 	if i%2 != 0 {
	// 		out <- i // Send 'i' to channel 'out'.
	// 	}
	// }
}
