package goroutine

import (
	"fmt"
	"testing"
	"time"
)

/*
同一个操作符 <- 既用于发送也用于接收
发送：ch <- "Washington"
接收：str = <- ch
*/
func TestChannelDemo(t *testing.T) {

	// 知识点一：使用了协程的 lambda 函数
	// 知识点二：使用isok通道判断协程是否执行完毕
	// 知识点三：使用defer自动关闭通道
	channelData()

	// 知识点：使用 for-range 语句会自动检测通道是否关闭
	sendChan := make(chan int)
	receiveChan := make(chan string)
	go processChannel(sendChan, receiveChan)

}

func sendData() chan string {
	// 缓冲区可接收任何类型元素，有缓冲区相当于异步
	// 带缓冲通道
	buf := 100
	ch := make(chan string, buf)
	// lambda 函数
	go func() {
		ch <- "Washington"
		ch <- "Tripoli"
		ch <- "London"
		ch <- "Beijing"
		ch <- "Tokyo"
		//通道可以被显式的关闭；尽管它们和文件不同：不必每次都关闭。
		//只有在当需要告诉接收者不会再提供新的值的时候，才需要关闭通道。
		//只有发送者需要关闭通道，接收者永远不会需要。
		close(ch)
	}()
	return ch
}

// 有方向的通道参数
// chan<-  	// 通道仅接收数据
// <-chan	// 通道仅发送数据
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

func channelData() {
	// 创建通道，协程在通道内进行数据交换
	// 无缓冲，一次只能包含一个元素
	isok := make(chan string)
	// 方法执行完后执行关闭通道
	defer close(isok)
	ch := sendData()
	for i := 0; i < 3; i++ {
		go getData(ch, isok)
	}
	// time.Sleep(5e9)
	for i := 0; i < 3; i++ {
		// 使用通道阻塞特性检测是否协程是否处理完成
		<-isok
	}
}

// 有方向的通道参数
// chan<-  	// 通道仅接收数据
// <-chan	// 通道仅发送数据
func processChannel(in <-chan int, out chan<- string) {
	//使用 for-range 语句会自动检测通道是否关闭
	for inValue := range in {
		result := string(rune(inValue)) /// processing inValue
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
