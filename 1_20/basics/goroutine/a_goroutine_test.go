package goroutine

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestGoroutineDemo(t *testing.T) {
	runtime.GOMAXPROCS(2) // 在2个核心CPU上并行运行
	fmt.Println("In main()")
	//使用go启动协程
	go longWait()
	go shortWait()
	fmt.Println("About to sleep in main()")
	// sleep works with a Duration in nanoseconds (ns) !
	time.Sleep(10 * 1e9)
	fmt.Println("At the end of main()")
	/* 结果：
	In main()
	About to sleep in main()
	Beginning shortWait()
	Beginning longWait()
	End of shortWait()
	End of longWait()
	At the end of main()
	*/
}

func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * 1e9) // sleep for 5 seconds
	fmt.Println("End of longWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * 1e9) // sleep for 2 seconds
	fmt.Println("End of shortWait()")
}
