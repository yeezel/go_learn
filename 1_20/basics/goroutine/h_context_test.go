package goroutine

import (
	"context"
	"fmt"
	"time"
)

type Options struct{ Interval time.Duration }

func reqTask(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done(): //等待cancel()执行并退出
			fmt.Println("stop", name)
			return
		default:
			fmt.Println(name, "send request")
			// 获取参数
			op := ctx.Value("options").(*Options)
			time.Sleep(op.Interval * time.Second)
		}
	}
}

// 并发控制子协程
func TestContextDemo() {
	//context.Backgroud() 创建根 Context，通常在 main 函数、初始化和测试代码中创建，作为顶层 Context。
	ctx, cancel := context.WithCancel(context.Background())

	// 控制子协程的最长执行时间，子协程在2秒后退出
	// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	// 控制子协程的最迟退出时间，这里设置具体的截至时间点，超过时间点将视为超时
	//ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))

	//给协程传递参数
	vCtx := context.WithValue(ctx, "options", &Options{1})

	go reqTask(vCtx, "worker1")
	go reqTask(vCtx, "worker2")
	time.Sleep(3 * time.Second)
	cancel() //退出子协程
	time.Sleep(3 * time.Second)
}
