/*
	WaitGroup 和信道(channel)是常见的 2 种并发控制的方式。
	如果并发启动了多个子协程，需要等待所有的子协程完成任务，WaitGroup 非常适合于这类场景

	使用锁的情景：
		访问共享数据结构中的缓存信息
		保存应用程序上下文和状态信息数据
*/

package goroutine

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func doTask(n int) {
	time.Sleep(time.Duration(n))
	fmt.Printf("Task %d Done\n", n)
	wg.Done()
}

func SyncWaitGroup() {
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go doTask(i + 1)
	}
	//wg.Wait() 会等待所有的子协程任务全部完成，所有子协程结束后，才会执行 wg.Wait() 后面的代码。
	wg.Wait()
	fmt.Println("All Task Done")
}
